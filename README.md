# Re:gret

<img src="cpanel/public/logo.png" width="200px">

**IaaS = Inquiry as a Service**

### 開発環境

```bash
# 初回のみ実行
# 実行後、.envに値を設定する
$ make deps
# 開発用サーバとDBの起動
# localhost:1323で立ち上がる
$ make dev
# DBマイグレーションの実行
$ make migrate/fresh
# シーダーの実行
$ make seed
# テスト
$ make test/inquiry
$ make test/category
```

### 本番環境

- 通常はGithub Actions経由で操作を行う
- 基本的に初回構築時のみ、サーバ内で操作する

```bash
$ make up -f 
$ make down
$ make migrate/fresh
```

- コンパネの環境構築は [こちら](./cpanel/README.md) を参照
- その他詳細に関しては各Makefileを参照

### 共通

```bash
$ make key
```
