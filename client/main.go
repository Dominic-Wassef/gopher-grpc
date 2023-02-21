package main

import (
	"context"
	"log"
	"time"

	pb "dominic.com/gopher-grpc/proto"

	"google.golang.org/grpc"
)

const (
	// Port for gRPC server to listen to
	ADDRESS = "localhost:50051"
)

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	todos := []TodoTask{
		{Name: "Code review", Description: "Review a new code feature", Done: false},
		{Name: "Make youtube view", Description: "Start go for GRPC", Done: false},
		{Name: "Go to gym", Description: "Leg day", Done: false},
		{Name: "Buy grocerties", Description: "Buy onion magos and bananas", Done: false},
		{Name: "Meet with mentor", Description: "Discuss blockers with my project", Done: false},
	}

	for _, todo := range todos {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`
		ID : %s
		Name : %s
		Description : %s
		Done : %v,
	`, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
	}
}
