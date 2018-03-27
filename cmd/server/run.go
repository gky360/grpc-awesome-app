package main

import (
	"os"

	"github.com/gky360/grpc-awesome-app/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		return 1
	}
	return 0
}
