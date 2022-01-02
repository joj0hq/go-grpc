package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jotaro-biz/go-remote-debug/pb"
	"google.golang.org/grpc"
)

func CreateTask(title string) {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewToDoServiceClient(conn)
	createTaskRequest := pb.CreateTaskRequest{
		Title: title,
	}

	reply, err := client.CreateTask(ctx, &createTaskRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	// レスポンスを標準出力に表示します。
	replyedTitle := reply.GetTitle()
	replyedCreateTime := reply.GetCreateTime()
	fmt.Println("***********************************")
	fmt.Printf("title: %s \n", replyedTitle)
	fmt.Printf("create time: %s \n", replyedCreateTime.String())
	fmt.Println("***********************************")
	fmt.Println()
}

func ShowList() {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewToDoServiceClient(conn)
	showListRequest := pb.ShowListRequest{}

	reply, err := client.ShowList(ctx, &showListRequest)
	if err != nil {
		log.Fatal("Request failed.")
		return
	}

	// レスポンスを標準出力に表示します。
	report := reply.GetTasks()
	if len(report) == 0 {
		fmt.Println("***********************************")
		fmt.Println("There are no tasks")
		fmt.Println("***********************************")
		fmt.Println()
		return
	}

	fmt.Println("***********************************")
	for k, v := range report {
		fmt.Println(k + 1)
		fmt.Printf("Title: %s \n", v.Title)
		fmt.Printf("Create time: %s \n", v.CreateTime.String())
		fmt.Println()
	}

	fmt.Printf("Amount: %d \n", reply.GetAmount())
	fmt.Println("***********************************")
	fmt.Println()
}
