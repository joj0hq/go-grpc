package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/jotaro-biz/go-remote-debug/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// DBを使わず構造体でデータを保持
type ToDoService struct {
	amount int64
	tasks  []*pb.Task
}

func NewToDoService() *ToDoService {
	return &ToDoService{
		amount: 0,
		tasks:  make([]*pb.Task, 0),
	}
}

func (s *ToDoService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	now := time.Now()
	task := pb.Task{
		Title: req.Title,
		CreateTime: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}
	s.tasks = append(s.tasks, &task)
	s.amount++

	return &pb.CreateTaskResponse{
		Title:      task.Title,
		CreateTime: task.CreateTime,
	}, nil
}

func (s *ToDoService) ShowList(ctx context.Context, req *pb.ShowListRequest) (*pb.ShowListResponse, error) {
	return &pb.ShowListResponse{
		Tasks:  s.tasks,
		Amount: s.amount,
	}, nil
}
