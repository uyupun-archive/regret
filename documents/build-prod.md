# 環境構築(本番環境)

## 1. APP_KEYの生成

- 管理用APIのアクセスに必要

```bash
# 実行後、.envに適切な値を設定する
$ cp .env.example .env
$ cd cpanel && cp .env.local.example .env.local
$ make key -f Makefile.prod
```

## 2. APIの環境構築

- 通常はGithub Actions経由で操作を行う
- 基本的に初回構築時のみ、サーバ内で操作する

```bash
# サーバとDBの起動
$ make up -f Makefile.prod
# サーバとDBの終了
$ make down -f Makefile.prod
# プロセスの確認
$ make ps
# DBマイグレーション
$ make migrate/up -f Makefile.prod
```

- その他のコマンドの詳細に関しては[Makefile](../Makefile)を参照

## 3. コンパネの環境構築

- 通常はGithub Actions経由で操作を行う
- 基本的に初回構築時のみ、サーバ内で操作する

```bash
$ cd cpanel
# 実行後、.env.productionに適切な値を設定する
$ cp .env.production.local.example .env.production.local
# サーバの起動
$ make up
# サーバの終了
$ make down
# プロセスの確認
$ make ps
```

- その他のコマンドの詳細に関しては[Makefile](../cpanel/Makefile)を参照
