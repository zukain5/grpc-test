# grpc-test
gRPC のおためし

# 概要
Person の情報を保持する簡単なサーバー。
gRPC のサーバーとして、以下の API を実行できる

- `GetFeature`: id から Person の情報を取得できる

# サーバーの実行方法
```shell
go run server/server.go
```

`localhost:50051` にサーバーが立ちます。
