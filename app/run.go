package app

import (
	"github.com/izumin5210-sandbox/todoapp-grapi/app/server"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewTodoServiceServer(),
		),
	)
	return s.Serve()
}
