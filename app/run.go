package app

import (
	"github.com/gky360/grpc-awesome-app/app/server"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewBookServiceServer(),
		),
	)
	return s.Serve()
}
