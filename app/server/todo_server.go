package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/izumin5210-sandbox/todoapp-grapi/api"
)

// TodoServiceServer is a composite interface of api_pb.TodoServiceServer and grapiserver.Server.
type TodoServiceServer interface {
	api_pb.TodoServiceServer
	grapiserver.Server
}

// NewTodoServiceServer creates a new TodoServiceServer instance.
func NewTodoServiceServer() TodoServiceServer {
	return &todoServiceServerImpl{}
}

type todoServiceServerImpl struct {
}

func (s *todoServiceServerImpl) ListTodos(ctx context.Context, req *api_pb.ListTodosRequest) (*api_pb.ListTodosResponse, error) {
	return &api_pb.ListTodosResponse{
		Todos: []*api_pb.Todo{
			&api_pb.Todo{TodoId: "1", Title: "Write Go 4 Advent Calendar at 2018/12/05", Done: true},
			&api_pb.Todo{TodoId: "2", Title: "Write Go 2 Advent Calendar at 2018/12/15", Done: true},
			&api_pb.Todo{TodoId: "3", Title: "Write Go 3 Advent Calendar at 2018/12/20", Done: true},
			&api_pb.Todo{TodoId: "4", Title: "Write Go 3 Advent Calendar at 2018/12/20", Done: false},
		},
	}, nil
}

func (s *todoServiceServerImpl) GetTodo(ctx context.Context, req *api_pb.GetTodoRequest) (*api_pb.Todo, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *todoServiceServerImpl) CreateTodo(ctx context.Context, req *api_pb.CreateTodoRequest) (*api_pb.Todo, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *todoServiceServerImpl) UpdateTodo(ctx context.Context, req *api_pb.UpdateTodoRequest) (*api_pb.Todo, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *todoServiceServerImpl) DeleteTodo(ctx context.Context, req *api_pb.DeleteTodoRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
