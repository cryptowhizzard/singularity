package run

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/dealpusher"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/urfave/cli/v2"
)

var DealPusherCmd = &cli.Command{
	Name:  "deal-pusher",
	Usage: "Start a deal pusher that monitors deal schedules and pushes deals to storage providers",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:    "deal-attempts",
			Usage:   "Number of times to attempt a deal before giving up",
			Aliases: []string{"d"},
			Value:   3,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		defer closer.Close()
		if err := model.AutoMigrate(db); err != nil {
			return err
		}

		lotusAPI := c.String("lotus-api")
		lotusToken := c.String("lotus-token")
		err = epochutil.Initialize(c.Context, lotusAPI, lotusToken)
		if err != nil {
			return err
		}

		dm, err := dealpusher.NewDealPusher(db, c.String("lotus-api"), c.String("lotus-token"), c.Uint("deal-attempts"))
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return service.StartServers(c.Context, dealpusher.Logger, dm)
	},
}