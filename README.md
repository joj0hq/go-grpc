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

## コンテナの起動

```bash
docker-compose up
```

## grpc_cliでの動作確認
command
```bash
grpc_cli ls localhost:50051
```

response
```bash
grpc.reflection.v1alpha.ServerReflection
todo_app.ToDoService
```

command
```bash
grpc_cli ls localhost:50051 todo_app.ToDoService -l
```

response
```bash
filename: proto/todo.proto
package: todo_app;
service ToDoService {
  rpc CreateTask(todo_app.CreateTaskRequest) returns (todo_app.CreateTaskResponse) {}
  rpc ShowList(todo_app.ShowListRequest) returns (todo_app.ShowListResponse) {}
}
```

command
```bash
grpc_cli call localhost:50051 todo_app.ToDoService.CreateTask 'title: "HELLO"'
```

response
```bash
connecting to localhost:50051
title: "HELLO"
create_time {
  seconds: 1641139300
  nanos: 49475000
}
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