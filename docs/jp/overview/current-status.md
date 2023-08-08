---
description: >-
  このプロジェクトは現在アクティブな開発中です。以下に、現在の機能リストとそのステータスを示します。
---

# 現状

このプロジェクトは現在アクティブな開発中です。以下に、現在の機能リストとそのステータスを示します。

| バッジ                                                       | 説明                                                     |
| ----------------------------------------------------------- | -------------------------------------------------------- |
| ![Stable](https://img.shields.io/badge/-Stable-brightgreen) | ステーブルであり、本番環境で使用する準備ができている機能     |
| ![Beta](https://img.shields.io/badge/-Beta-blue)            | ベータ版であり、まだバグを含んでいる可能性がある機能         |
| ![Alpha](https://img.shields.io/badge/-Alpha-orange)        | アルファ版であり、本番環境で使用するべきではない機能         |
| ![WIP](https://img.shields.io/badge/-WIP-yellow)            | 現在作業中であり、使用できない機能                         |
| ![Planned](https://img.shields.io/badge/-Planned-lightgrey) | 予定されているがまだ実装されていない機能                    |

<table><thead><tr><th width="158">カテゴリ</th><th width="186">機能</th><th width="106">ステータス</th><th>説明</th></tr></thead><tbody><tr><td>データソース</td><td>ファイルシステム</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>ローカルファイルシステムでデータを準備するためのサポート</td></tr><tr><td>データソース</td><td>その他のリモート</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>rcloneによってバックアップされたその他のリモートシステムからデータを準備するためのサポート</td></tr><tr><td>データ準備</td><td>データセットの作成</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>データセットを作成するためのCLIツール</td></tr><tr><td>データ準備</td><td>データソースの追加</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>既存のデータセットにデータソースを追加するためのCLIツール</td></tr><tr><td>データ準備</td><td>インラインプリパレーション</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>インラインプリパレーションのサポート。CARファイルのエクスポートは不要です</td></tr><tr><td>データ準備</td><td>アップロードAPI</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>APIを介してファイルを手動でアップロードするためのサポート</td></tr><tr><td>データ準備</td><td>プッシュAPI</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>APIを介して新しいアイテムをキューに追加するためのサポート</td></tr><tr><td>データ準備</td><td>DAG再生成</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>データセット全体のルートCIDを更新するためにUnixfs DAGを再生成する機能</td></tr><tr><td>データ準備</td><td>基本暗号化</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>非対称キーを使用した基本的な暗号化のサポート</td></tr><tr><td>データ準備</td><td>カスタム暗号化</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>ユーザーが暗号化ツールを提供するカスタム暗号化のサポート</td></tr><tr><td>コンテンツプロバイダー</td><td>HTTPピース</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>HTTPピースのダウンロード（CARファイルのダウンロード）のサポート</td></tr><tr><td>コンテンツプロバイダー</td><td>IPFSゲートウェイ</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>IPFS Gatway準拠のデータ取得のサポート</td></tr><tr><td>コンテンツプロバイダー</td><td>Bitswap</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>Bitswap取得（IPFS間の相互運用）のサポート</td></tr><tr><td>コンテンツプロバイダー</td><td>Graphsync</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>Graphsync取得のサポート</td></tr><tr><td>コンテンツプロバイダー</td><td>メタデータAPI</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>元のデータオーナーからのCARファイル配布の許可</td></tr><tr><td>コンテンツプロバイダー</td><td>ダウンロードクライアント</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>メタデータAPIの助けを借りて、元のデータオーナーからCARファイルをダウンロードするサポート</td></tr><tr><td>ディールメイキング</td><td>ディールトラッキング</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>ディールのステータスを追跡する機能</td></tr><tr><td>ディールメイキング</td><td>ディールスケジューラ</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td><a href="https://github.com/tech-greedy/singularity/tree/main#deal-replication">js-singularity</a>ディールレプリケーションと同等の機能</td></tr><tr><td>ディールメイキング</td><td>Spade API</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>ストレージプロバイダーが自己提案ディールをするためのSpade APIの互換性のサポート</td></tr><tr><td>ディールメイキング</td><td>ウォレット管理</td><td><img src="https://img.shields.io/badge/-Alpha-orange" alt="Alpha"></td><td>ウォレット管理のサポート</td></tr><tr><td>ディールメイキング</td><td>リモートサインカー</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>リモートサインカーのサポート</td></tr><tr><td>ユーティリティ</td><td>ベンチマーク</td><td><img src="https://img.shields.io/badge/-Beta-blue" alt="Beta"></td><td>データ準備のベンチマークのサポート</td></tr><tr><td>ユーティリティ</td><td>メトリクス収集</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>メトリクスの収集のサポート</td></tr><tr><td>ユーティリティ</td><td>モニタリング</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>データ準備とディールメイキングのモニタリングのサポート</td></tr><tr><td>ダッシュボード</td><td>データセットエクスプローラ</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>データセットをフォルダごとに探索するためのサポート</td></tr><tr><td>ダッシュボード</td><td>データセットのダウンロード</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>ブラウザで直接データセットをダウンロードするためのサポート</td></tr><tr><td>ダッシュボード</td><td>ディールエクスプローラ</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>ディール提案の探索のサポート</td></tr><tr><td>ダッシュボード</td><td>ピースエクスプローラ</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>ピースCIDで探索し、配布状況をチェックするためのサポート</td></tr><tr><td>ダッシュボード</td><td>提供者ビュー</td><td><img src="https://img.shields.io/badge/-Planned-lightgrey" alt="Planned"></td><td>提供者がディールを消費している速度を確認するためのサポート</td></tr></tbody></table>