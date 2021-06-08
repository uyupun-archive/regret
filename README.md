# Re:gret

<img src="cpanel/public/logo.png" width="200px">

**IaaS = Inquiry as a Service**

### 環境構築（API）

- プロジェクトルートで実行

```bash
$ make init
$ make db
$ make api  # localhost:1323
```

### 環境構築（コンパネ）

- `cpanel` ディレクトリ内で実行

```bash
$ cp .env.example .env
$ yarn dev  # localhost:3000
```

### 環境構築（共通）

- プロジェクトルートで実行

```bash
$ make key
```
