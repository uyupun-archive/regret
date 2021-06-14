# Re:gret コントロールパネル

### 開発環境

```bash
# 開発用サーバの起動
# localhost:3000
$ yarn dev
```

### 本番環境

- 通常はGithub Actions経由で操作を行う
- 基本的に初回構築時のみ、サーバ内で操作する

```bash
$ cp .env.production.local.example .env.production.local
$ make up
$ make down
$ make ps
```
