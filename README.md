# Re:gret

<img src="cpanel/public/logo.png" width="200px">

**IaaS = Inquiry as a Service**

## 構築手順

### 開発環境

1. APP_KEYの生成

- 管理用APIのアクセスに必要

```bash
$ cd cpanel && cp .env.local.example .env.local
$ make key
```

1. APIの環境構築

- 開発環境にGo 1.16+がインストールされている前提

```bash
# 初回のみ実行
# 実行後、.envに適切な値を設定する
$ make deps
# 開発用サーバとDBの起動
# localhost:1323で立ち上がる
$ make dev
# DBマイグレーションの実行
$ make migrate/fresh
# シーダーの実行
$ make seed
# 外部公開APIの疎通テスト
$ make test/inquiry
$ make test/category
```

- その他のコマンドに関してはMakefileを参照

1. コンパネの環境構築

- コンパネの環境構築は [こちら](./cpanel/README.md) を参照

### 本番環境

1. APP_KEYの生成

- 管理用APIのアクセスに必要

```bash
$ cd cpanel && cp .env.local.example .env.local
$ make key
```

1. APIの環境構築

- 通常はGithub Actions経由で操作を行う
- 基本的に初回構築時のみ、サーバ内で操作する

```bash
# サーバとDBの起動
$ make up -f Makefile.prod
# サーバとDBの終了
$ make down -f Makefile.prod
# DBマイグレーション
$ make migrate/fresh -f Makefile.prod
```

- その他のコマンドに関してはMakefileを参照

1. コンパネの環境構築

- コンパネの環境構築は [こちら](./cpanel/README.md) を参照
