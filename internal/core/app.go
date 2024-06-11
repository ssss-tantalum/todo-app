package core

import (
	"database/sql"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type App struct {
	router *echo.Echo
	cfg    *Config

	dbOnce sync.Once
	db     *bun.DB
}

func NewApp(cfg *Config) *App {
	app := &App{
		cfg: cfg,
	}
	app.initRouter()

	return app
}

func (app *App) IsDebug() bool {
	return app.cfg.Debug
}

func (app *App) Cfg() *Config {
	return app.cfg
}

func (app *App) Router() *echo.Echo {
	return app.router
}

func (app *App) DB() *bun.DB {
	app.dbOnce.Do(func() {
		sqldb, err := sql.Open(sqliteshim.ShimName, app.cfg.DBdsn)
		if err != nil {
			panic(err)
		}

		db := bun.NewDB(sqldb, sqlitedialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithEnabled(app.IsDebug()),
			bundebug.FromEnv(""),
		))

		app.db = db
	})
	return app.db
}
