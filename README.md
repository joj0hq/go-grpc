# go-grpc
This is grpc server for golang.

## protocのインストール

```bash
brew install protoc
```

## Golang用のプラグインのインストール

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
```

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## .protoファイルから自動生成

```bash
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/todo.proto
```

## ローカルでgrpcを使うためのインストール

```bash
brew tap grpc/grpc
```

```bash
brew install grpc
```

## ドキュメント生成用のインストール

```bash
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
```

## ドキュメントの自動生成

```bash
protoc --doc_out=html,index.html:./doc proto/*.proto
```

## 参考資料
- https://grpc.io/docs/languages/go