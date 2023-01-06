フロントエンド: vue  
バックエンド: golang

## docker環境の立ち上げ方
```
docker-compose up -d
```

仮想環境  
+ vue: フロント描画用  
  + ソースコード: appディレクトリ  
+ go: apiサーバ  
  + ソースコード: assetsディレクトリ  
+ mysql: DBサーバ  
  + マイグレーションファイル: dbディレクトリ 

## DBマイグレーション
```
cd db
./migrate -database="mysql://root:password@tcp(mysql:3306)/sample" -path=./migrations/ up
```
## vueコンパイル
```
docker-compose exec vue sh
npm run serve
```
## サーバ立ち上げ
```
docker-compose exec go bash
go run main.go
```

## ログイン画面  
http://localhost:8080/login  
usersテーブルの中で、email・passwordの一致するユーザでログインできる。
