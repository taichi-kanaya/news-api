# news-api

## 概要
さまざまなサイトから最新のニュースを取得するAPIです。  
2023年5月現在、下記APIを利用しています。
1. News API  
https://newsapi.org/
1. Reddit API(※coming soon)

1. GoogleNewsAPI(※coming soon)

1. BingNewsSearchAPI(※coming soon)

---

## 開発環境構築
### 【前提】
Go 1.20 をインストールしてください。

### 【環境変数】
下記環境変数の設定が必要です。  

|環境変数名|説明|
|-|-|
|APP_ENV|dev or prod|
|NEWS_API_KEY|News APIのAPIキー|
|SENTRY_DSN|SentryのDSN（= エラーログ送信先URL）|

### 【テスト実行】
ルートディレクトリで以下を実行してください。
```bash
go test ./...
```

### 【ローカル起動】
ルートディレクトリで以下を実行してください。
```bash
go run main.go
```

---
## 補足事項
### 【OpenAPI について】
OpenAPI Specification 2.0を導入しております。  
ローカルで起動後、以下のURLから確認できます。  
http://localhost:8080/swagger/index.html

### 【godoc について】
ルートディレクトリで以下を実行してください。
```bash
godoc -http=:6060
```
これにより、  
http://localhost:6060/pkg/news-api/  
でドキュメントを表示させることができます。

