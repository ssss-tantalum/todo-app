package main

import (
	"github.com/ssss-tantalum/gotth-stack/internal/core"
	"github.com/ssss-tantalum/gotth-stack/internal/server"
)

func main() {
	cfg := core.NewConfig()
	app := core.NewApp(cfg)

	server.Serve(app)
}
