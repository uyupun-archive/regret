# 管理API仕様

- ベースパス: `/api/v0/admin`
- BearerトークンにはAPP_KEYを含める

### サービス一覧取得

- HTTPメソッド: `GET`
- エンドポイント: `/service`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{}
```
- レスポンスボディ:
```json
[
  {
    "id": 1,
    "name": "hoge",
    "name_ja": "ほげ",
    "description": "hogehogehogehoge",
    "access_token": "xxxxxxxxxx"
  },
  ...
]
```

### サービス追加

- HTTPメソッド: `POST`
- エンドポイント: `/service`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "name": "hoge",
  "name_ja": "ほげ",
  "description": "hogehogehogehoge"
}
```
- レスポンスボディ:
```json
{}
```

### サービス編集

- HTTPメソッド: `PATCH`
- エンドポイント: `/service`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "id": 1,
  "name": "hoge",
  "name_ja": "ほげ",
  "description": "hogehogehogehoge",
  "access_token": "xxxxxxxxxx"
}
```
- レスポンスボディ:
```json
{}
```

### サービス削除

- HTTPメソッド: `DELETE`
- エンドポイント: `/service`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "id": 1,
  "name": "hoge",
  "name_ja": "ほげ",
  "description": "hogehogehogehoge",
  "access_token": "xxxxxxxxxx"
}
```
- レスポンスボディ:
```json
{}
```

### カテゴリ一覧取得

- HTTPメソッド: `GET`
- エンドポイント: `/category`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "service_id": 1
}
```
- レスポンスボディ:
```json
[
  {
    "id": 1,
    "service_id": 1,
    "name": "piyo",
    "name_ja": "ぴよ"
  },
  ...
]
```

### カテゴリ追加

- HTTPメソッド: `POST`
- エンドポイント: `/category`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "id": 1,
  "service_id": 1,
  "name": "piyo",
  "name_ja": "ぴよ"
}
```
- レスポンスボディ:
```json
{}
```

### カテゴリ編集

- HTTPメソッド: `PATCH`
- エンドポイント: `/category`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "id": 1,
  "service_id": 1,
  "name": "piyo",
  "name_ja": "ぴよ"
}
```
- レスポンスボディ:
```json
{}
```

### カテゴリ削除

- HTTPメソッド: `DELETE`
- エンドポイント: `/category`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "id": 1,
  "service_id": 1,
  "name": "piyo",
  "name_ja": "ぴよ"
}
```
- レスポンスボディ:
```json
{}
```
