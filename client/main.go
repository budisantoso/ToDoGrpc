package main

import (
	"context"
	"flag"
	"log"
	pb "projects/Grpc/ToDoList/api"

	"google.golang.org/grpc"
)

func main() {
	//1. Getting server ip that can be connected to
	//2. Do grpc dial
	//3. Register grpc dial connection to client server
	//4. Call grpc api

	backend := flag.String("b", "localhost:8082", "server to connect to")
	flag.Parse()
	log.Printf("connect to server %v", *backend)

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not dial to server %v", err)
	}

	client := pb.NewTodoServiceClient(conn)

	//create todo
	todonew := &pb.CreateTodoRequest{
		Name:   "cooking",
		Desc:   "sambal",
		Status: false,
	}
	todores, err := client.Create(context.Background(), todonew)

	if err != nil {
		log.Fatalf("can not create new task %v", err)
	}
	log.Printf("new task is created with id: %v", todores.Tid)

	//show created todo
	listtosearch := &pb.ListTodoRequest{
		Tid: todores.Tid,
	}
	todolist, err := client.List(context.Background(), listtosearch)
	if err != nil {
		log.Fatalf("can not get task with id %v and error is %v", todores.Tid, err)
	}
	log.Printf("task with id %v is %v", todores.Tid, todolist)
}
