package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jotaro-biz/go-remote-debug/pb"
	"github.com/jotaro-biz/go-remote-debug/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterToDoServiceServer(server, service.NewToDoService())

	// サーバーリフレクションを有効にして、シリアライズなしでgrpc_cliで動作確認ができる
	reflection.Register(server)
	server.Serve(listenPort)
}
