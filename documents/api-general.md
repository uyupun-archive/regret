# 一般API仕様

- 外部サービスから使用するAPI
- ベースパス: `/api/v0`
- Bearerトークンにはアクセストークンを含める

### カテゴリ一覧取得

- HTTPメソッド: `GET`
- エンドポイント: `/category`
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
    "service_id": 1,
    "name": "piyo",
    "name_ja": "ぴよ"
  },
  ...
]
```

### 問い合わせ

- HTTPメソッド: `POST`
- エンドポイント: `/category`
- リクエストヘッダ: `Authorization: Bearer xxxxxxxxxx`
- リクエストボディ:
```json
{
  "subject": "ふが",
  "email": "uyupun@gmail.com",
  "category_id": 1,
  "text": "ふがふが"
}
```
- レスポンスボディ:
```json
{}
```
