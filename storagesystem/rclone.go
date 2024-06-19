package storagesystem

import (
    "bufio"
    "bytes"
    "context"
    "io"
    "os/exec"
    "sort"
    "strings"
    "sync"
    "time"

    "github.com/cockroachdb/errors"
    "github.com/data-preservation-programs/singularity/model"
    "github.com/gammazero/workerpool"
    "github.com/ipfs/go-log/v2"
    "github.com/rclone/rclone/fs"
    "github.com/rclone/rclone/fs/config/configmap"
    "github.com/rclone/rclone/fs/object"
    "go.uber.org/zap"
)

var logger = log.Logger("storage")

var _ Handler = &RCloneHandler{}

var ErrGetUsageNotSupported = errors.New("The backend does not support getting usage quota")
var ErrBackendNotSupported = errors.New("This backend is not supported")
var ErrMoveNotSupported = errors.New("The backend does not support moving files")

type RCloneHandler struct {
    name                    string
    fs                      fs.Fs
    fsNoHead                fs.Fs
    retryMaxCount           int
    retryDelay              time.Duration
    retryBackoff            time.Duration
    retryBackoffExponential float64
    scanConcurrency         int
    logger                  *zap.Logger
}

func (h RCloneHandler) Name() string {
    return h.name
}

func (h RCloneHandler) Write(ctx context.Context, path string, in io.Reader) (fs.Object, error) {
    objInfo := object.NewStaticObjectInfo(path, time.Now(), -1, true, nil, nil)
    if strings.HasSuffix(path, ".car") {
        objInfo = objInfo.WithMimeType("application/vnd.ipfs.car")
    }
    return h.fs.Put(ctx, in, objInfo)
}

func (h RCloneHandler) Move(ctx context.Context, from fs.Object, to string) (fs.Object, error) {
    if h.fs.Features().Move != nil {
        return h.fs.Features().Move(ctx, from, to)
    }
    return nil, errors.Wrapf(ErrMoveNotSupported, "backend: %s", h.fs.String())
}

func (h RCloneHandler) Remove(ctx context.Context, obj fs.Object) error {
    return obj.Remove(ctx)
}

func (h RCloneHandler) About(ctx context.Context) (*fs.Usage, error) {
    h.logger.Debug("About: getting usage", zap.String("type", h.fs.Name()))
    if h.fs.Features().About != nil {
        return h.fs.Features().About(ctx)
    }

    return nil, errors.Wrapf(ErrGetUsageNotSupported, "backend: %s", h.fs.String())
}

func (h RCloneHandler) List(ctx context.Context, path string) ([]fs.DirEntry, error) {
    h.logger.Debug("List: listing path", zap.String("type", h.fs.Name()), zap.String("root", h.fs.Root()), zap.String("path", path))
    return h.fs.List(ctx, path)
}

func (h RCloneHandler) Command(ctx context.Context, args ...string) *exec.Cmd {
    cmd := exec.CommandContext(ctx, "rclone", args...)
    return cmd
}

func (h RCloneHandler) processLine(ctx context.Context, path string, text string) (Entry, error) {
    // Your implementation here
    // Placeholder return
    return Entry{}, nil
}

func (h RCloneHandler) scan(ctx context.Context, path string, ch chan<- Entry, wp *workerpool.WorkerPool, wg *sync.WaitGroup) {
	defer wg.Done()

	if ctx.Err() != nil {
		return
	}
	h.logger.Info("Scan: listing path", zap.String("type", h.fs.String()), zap.String("path", path))

	// Execute the rclone command to list files
	cmd := h.Command(ctx, "lsjson", "-R", "--hash", path)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		h.logger.Error("failed to create stdout pipe", zap.Error(err))
		return
	}

	if err := cmd.Start(); err != nil {
		h.logger.Error("failed to start command", zap.Error(err))
		return
	}

	// Use bufio.Scanner for buffered reading
	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(make([]byte, 0, 128*1024), 2*1024*1024) // Set a larger buffer size

	var entries []fs.DirEntry
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := h.processLine(ctx, path, line)
		if err != nil {
			h.logger.Error("failed to process line", zap.String("line", line), zap.Error(err))
			continue
		}
		if entry.Dir != nil {
			entries = append(entries, entry.Dir)
		}
		if entry.Info != nil {
			entries = append(entries, entry.Info)
		}
	}

	if err := scanner.Err(); err != nil {
		h.logger.Error("error reading stdout", zap.Error(err))
	}

	if err := cmd.Wait(); err != nil {
		h.logger.Error("command finished with error", zap.Error(err))
	}

	// Process and sort the entries
	sort.Slice(entries, func(i, j int) bool {
		return strings.Compare(entries[i].Remote(), entries[j].Remote()) < 0
	})

	var subCount int
	for _, entry := range entries {
		switch v := entry.(type) {
		case fs.Directory:
			select {
			case <-ctx.Done():
				return
			case ch <- Entry{Dir: v}:
			}

			subPath := v.Remote()
			wg.Add(1)
			wp.Submit(func() {
				h.scan(ctx, subPath, ch, wp, wg)
			})
			subCount++
		case fs.Object:
			select {
			case <-ctx.Done():
				return
			case ch <- Entry{Info: v}:
			}
		}
	}

	h.logger.Debug("Scan: finished listing path", zap.Int("remaining paths to list", subCount))
}


func (h RCloneHandler) Scan(ctx context.Context, path string) <-chan Entry {
    ch := make(chan Entry, h.scanConcurrency)
    go func() {
        var wg sync.WaitGroup
        wp := workerpool.New(h.scanConcurrency)
        wg.Add(1)
        wp.Submit(func() {
            h.scan(ctx, path, ch, wp, &wg)
        })
        wg.Wait() // OK to wait while child scans continue adding to wg
        wp.StopWait()
        close(ch)
    }()
    return ch
}

func (h RCloneHandler) Check(ctx context.Context, path string) (fs.DirEntry, error) {
    h.logger.Debug("Check: checking path", zap.String("type", h.fs.Name()), zap.String("root", h.fs.Root()), zap.String("path", path))
    return h.fs.NewObject(ctx, path)
}

type readCloser struct {
    io.Reader
    io.Closer
}

type readerWithRetry struct {
    ctx                     context.Context
    object                  fs.Object
    reader                  io.ReadCloser
    offset                  int64
    retryDelay              time.Duration
    retryBackoff            time.Duration
    retryCountMax           int
    retryCount              int
    retryBackoffExponential float64
}

func (r *readerWithRetry) Close() error {
    if r.reader != nil {
        return r.reader.Close()
    }
    return nil
}

func (r *readerWithRetry) Read(p []byte) (int, error) {
    if r.ctx.Err() != nil {
        return 0, r.ctx.Err()
    }
    n, err := r.reader.Read(p)
    r.offset += int64(n)
    //nolint:errorlint
    if err == io.EOF || err == nil {
        return n, err
    }

    if r.retryCount >= r.retryCountMax {
        return n, err
    }

    // error is not EOF
    logger.Warnf("Read error: %s, retrying after %s", err, r.retryDelay)
    select {
    case <-r.ctx.Done():
        return n, errors.Join(err, r.ctx.Err())
    case <-time.After(r.retryDelay):
    }
    r.retryCount += 1
    r.retryDelay = time.Duration(float64(r.retryDelay) * r.retryBackoffExponential)
    r.retryDelay += r.retryBackoff
    r.reader.Close()
    var err2 error
    r.reader, err2 = r.object.Open(r.ctx, &fs.SeekOption{Offset: r.offset})
    if err2 != nil {
        return n, errors.Join(err, err2)
    }
    return n, nil
}

func (h RCloneHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
    h.logger.Debug("Read: reading path", zap.String("type", h.fs.Name()), zap.String("root", h.fs.Root()), zap.String("path", path), zap.Int64("offset", offset), zap.Int64("length", length))
    if length == 0 {
        object, err := h.fs.NewObject(ctx, path)
        if err != nil {
            return nil, nil, errors.Wrapf(err, "failed to open object %s", path)
        }
        return io.NopCloser(bytes.NewReader(nil)), object, nil
    }
    object, err := h.fsNoHead.NewObject(ctx, path)
    if err != nil {
        return nil, nil, errors.Wrapf(err, "failed to open object %s", path)
    }
    option := &fs.SeekOption{Offset: offset}
    reader, err := object.Open(ctx, option)
    readerWithRetry := &readerWithRetry{
        ctx:                     ctx,
        object:                  object,
        reader:                  reader,
        offset:                  offset,
        retryDelay:              h.retryDelay,
        retryBackoff:            h.retryBackoff,
        retryCountMax:           h.retryMaxCount,
        retryBackoffExponential: h.retryBackoffExponential,
    }
    if length < 0 {
        return readerWithRetry, object, errors.WithStack(err)
    }
    return readCloser{
        Reader: io.LimitReader(readerWithRetry, length),
        Closer: readerWithRetry,
    }, object, errors.WithStack(err)
}

func NewRCloneHandler(ctx context.Context, s model.Storage) (*RCloneHandler, error) {
    _, ok := BackendMap[s.Type]
    registry, err := fs.Find(s.Type)
    if !ok || err != nil {
        return nil, errors.Wrapf(ErrBackendNotSupported, "type: %s", s.Type)
    }

    ctx, _ = fs.AddConfig(ctx)
    config := fs.GetConfig(ctx)
    overrideConfig(config, s)

    noHeadObjectConfig := make(map[string]string)
    headObjectConfig := make(map[string]string)
    for k, v := range s.Config {
        noHeadObjectConfig[k] = v
        headObjectConfig[k] = v
    }
    noHeadObjectConfig["no_head_object"] = "true"
    headObjectConfig["no_head_object"] = "false"

    noHeadFS, err := registry.NewFs(ctx, s.Type, s.Path, configmap.Simple(noHeadObjectConfig))
    if err != nil {
        return nil, errors.Wrapf(err, "failed to create RClone backend %s: %s", s.Type, s.Path)
    }

    headFS, err := registry.NewFs(ctx, s.Type, s.Path, configmap.Simple(headObjectConfig))
    if err != nil {
        return nil, errors.Wrapf(err, "failed to create RClone backend %s: %s", s.Type, s.Path)
    }

    scanConcurrency := 1
    if s.ClientConfig.ScanConcurrency != nil {
        scanConcurrency = *s.ClientConfig.ScanConcurrency
    }

    handler := &RCloneHandler{
        name:                    s.Name,
        fs:                      headFS,
        fsNoHead:                noHeadFS,
        retryMaxCount:           10,
        retryDelay:              time.Second,
        retryBackoff:            time.Second,
        retryBackoffExponential: 1.0,
        scanConcurrency:         scanConcurrency,
        logger:                  zap.NewExample(), // Initialize the logger here
    }

    if s.ClientConfig.RetryMaxCount != nil {
        handler.retryMaxCount = *s.ClientConfig.RetryMaxCount
    }
    if s.ClientConfig.RetryDelay != nil {
        handler.retryDelay = *s.ClientConfig.RetryDelay
    }
    if s.ClientConfig.RetryBackoff != nil {
        handler.retryBackoff = *s.ClientConfig.RetryBackoff
    }
    if s.ClientConfig.RetryBackoffExponential != nil {
        handler.retryBackoffExponential = *s.ClientConfig.RetryBackoffExponential
    }

    return handler, nil
}

func overrideConfig(config *fs.ConfigInfo, s model.Storage) {
    config.UseServerModTime = true
    if s.ClientConfig.ConnectTimeout != nil {
        config.ConnectTimeout = *s.ClientConfig.ConnectTimeout
    }
    if s.ClientConfig.Timeout != nil {
        config.Timeout = *s.ClientConfig.Timeout
    }
    if s.ClientConfig.ExpectContinueTimeout != nil {
        config.ExpectContinueTimeout = *s.ClientConfig.ExpectContinueTimeout
    }
    if s.ClientConfig.InsecureSkipVerify != nil {
        config.InsecureSkipVerify = true
    }
    if s.ClientConfig.NoGzip != nil {
        config.NoGzip = true
    }
    if s.ClientConfig.UserAgent != nil {
        config.UserAgent = *s.ClientConfig.UserAgent
    }
    if len(s.ClientConfig.CaCert) > 0 {
        config.CaCert = s.ClientConfig.CaCert
    }
    if s.ClientConfig.ClientCert != nil {
        config.ClientCert = *s.ClientConfig.ClientCert
    }
    if s.ClientConfig.ClientKey != nil {
        config.ClientKey = *s.ClientConfig.ClientKey
    }
    if len(s.ClientConfig.Headers) > 0 {
        for k, v := range s.ClientConfig.Headers {
            config.Headers = append(config.Headers, &fs.HTTPOption{
                Key:   k,
                Value: v,
            })
        }
    }
    if s.ClientConfig.DisableHTTP2 != nil {
        config.DisableHTTP2 = *s.ClientConfig.DisableHTTP2
    }
    if s.ClientConfig.DisableHTTPKeepAlives != nil {
        config.DisableHTTPKeepAlives = *s.ClientConfig.DisableHTTPKeepAlives
    }
    if s.ClientConfig.UseServerModTime != nil {
        config.UseServerModTime = *s.ClientConfig.UseServerModTime
    }
    if s.ClientConfig.LowLevelRetries != nil {
        config.LowLevelRetries = *s.ClientConfig.LowLevelRetries
    }
}
