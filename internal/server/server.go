package server

import (
	"fmt"

	"github.com/ssss-tantalum/gotth-stack/internal/core"
)

func Serve(app *core.App) {
	initRoutes(app)

	port := fmt.Sprintf(":%d", app.Cfg().HTTPPort)
	app.Router().Start(port)
}
