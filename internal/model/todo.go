package model

import "github.com/uptrace/bun"

type Todo struct {
	bun.BaseModel `bun:"table:todos"`

	ID      string `bun:"id,pk"`
	Content string `bun:"content,notnull"`
}
