package component

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetMainDBConnection() *mongo.Database
}

type appCtx struct {
	db *mongo.Database
}

func NewAppContext(db *mongo.Database) *appCtx {
	return &appCtx{
		db: db,
	}
}

func (ctx *appCtx) GetMainDBConnection() *mongo.Database {
	return ctx.db
}
