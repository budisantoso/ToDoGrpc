package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net"
	pb "projects/Grpc/ToDoList/api"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	//1. Create tcp connection
	//2. Create new grpc server
	//3. Register service to grpc
	//4. Serve grpc through tcp

	port := flag.Int("p", 8082, "port to listen to")
	flag.Parse()
	logrus.Printf("connected to port %d", *port)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("can not connect to tcp. err: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &todoServerService{
		TodoStorage: make(map[string]*pb.CreateTodoRequest),
	})
	if e := s.Serve(conn); e != nil {
		logrus.Fatalf("can't register server %v", e)
	}
}

type todoServerService struct {
	TodoStorage map[string]*pb.CreateTodoRequest
}

func (t *todoServerService) Create(c context.Context, todo *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	id := rand.Int()
	todoid := string(id)

	t.TodoStorage[todoid] = todo
	res := &pb.CreateTodoResponse{Tid: todoid}
	return res, nil
}

func (t *todoServerService) List(c context.Context, todo *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	todoRes := t.TodoStorage[todo.Tid]
	res := &pb.ListTodoResponse{
		Name:   todoRes.Name,
		Desc:   todoRes.Desc,
		Status: todoRes.Status,
	}
	return res, nil
}
