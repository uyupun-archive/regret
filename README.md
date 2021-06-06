# Re:gret

<img src="cpanel/public/logo.png" width="200px">

問い合わせデータの集約・通知システム

### 環境構築（API）

```bash
$ make init
$ make db
$ make api  # localhost:1323
```

### 環境構築（コンパネ）

```bash
$ cd cpanel
$ cp .env.local.example .env.local
$ yarn dev  # localhost:3000
```
