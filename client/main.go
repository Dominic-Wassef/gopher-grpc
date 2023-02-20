package main

import (
	"context"
	"log"
	"net"

	pb "dominic.com/gopher-grpc/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Recieved: %v", in.GetName())
}
