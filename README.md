# go-training

## migration
sql-migrate up  
sql-migrate down  

# 設計案

## アカウント作成画面
`/new_account POST`
JSON
```
{
  user_id: string,
  password: string,
  email_address: string
}
```

内部で一意なトークンを生成→ハッシュ化してメール認証専用のテーブルに保存
expires_atはunix timeで数時間後に設定する
```
table: email_activation
  token: string
  expires_at: int
```

ユーザーIDとパスワードのハッシュ値を格納したテーブルへの保存もこの時点で行う。ただし、まだ未認証なのでactivated: falseとする

```
table: user
  user_id: string
  email_address: string
  password: string
  activated: bool
```

そしてemail addressに認証メールを送る。
生トークンをURLに付与しておく。

## メール認証

`/new_account/validate?token=hoge GET`

`?token=hoge` のhogeをハッシュにかけてemail_activationで検索
→ある、かつexpires_atの期限内ならばアカウントを作成し、ログイン画面にリダイレクト→DBからそのレコードを消す
→あるが期限切れなら再認証画面へリダイレクトさせる→DBからそのレコードを消す
→DBになければ404 Not Found

## ログイン処理

`/login POST`

1. user_id
2. password

の2つを受け取る

`JSON` 

```
{
  user_id: string, 
  password: string
}
```

passwordをハッシュにかけてuser_id, hashed passwordをキーにuser tableを検索
→existsなら200, 以下に続く
→not_existsなら404を返す

## アクセストークンの発行
existsなら、
`/authorizations GET` 
にリダイレクトさせて、
tokenとrefresh_tokenを生成し、ユーザーに返却する

```
{
  access_token: string,
  refresh_token: string,
}
```

そしてトークンをhash化して以下のテーブルに保存する
expires_atは、access_tokenは短め（数時間）refresh_tokenは長く（数カ月）取る

```
credential
  user_id: string
  access_token: string
  access_token_expires_at: int
  refresh_token: string
  refresh_tokne_expires_at: int
```

