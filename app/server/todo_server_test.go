package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/izumin5210-sandbox/todoapp-grapi/api"
)

func Test_TodoServiceServer_ListTodos(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.ListTodosRequest{}

	resp, err := svr.ListTodos(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_GetTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.GetTodoRequest{}

	resp, err := svr.GetTodo(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_CreateTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateTodoRequest{}

	resp, err := svr.CreateTodo(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_UpdateTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateTodoRequest{}

	resp, err := svr.UpdateTodo(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TodoServiceServer_DeleteTodo(t *testing.T) {
	svr := NewTodoServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteTodoRequest{}

	resp, err := svr.DeleteTodo(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
