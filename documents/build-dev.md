# 環境構築(開発環境)

## 1. APP_KEYの生成

- 管理用APIのアクセスに必要

```bash
# 実行後、.envに適切な値を設定する
$ cp .env.example .env
$ cd cpanel && cp .env.local.example .env.local
$ make key
```

## 2. APIの環境構築

- 開発環境にGo 1.16+がインストールされている前提

```bash
# 依存パッケージのインストール
# 初回のみ実行
$ make deps
# 開発用のサーバとDBの起動
# localhost:1323で立ち上がる
$ make up
# DBの終了
$ make down
# プロセスの確認
$ make ps
# DBマイグレーションの実行
$ make migrate/fresh
# シーダーの実行
$ make seed
# 外部公開APIの疎通テスト
$ make test/inquiry
$ make test/category
# 本番環境で使用するスクリプトのビルド
$ make key/build
```

- その他のコマンドの詳細に関しては[Makefile](../Makefile)を参照

## 3. コンパネの環境構築

```bash
$ cd cpanel
# 依存パッケージのインストール
$ yarn install
# 開発用サーバの起動
# localhost:3000で立ち上がる
$ yarn dev
```

- その他のコマンドの詳細に関しては[Makefile](../cpanel/Makefile)を参照
