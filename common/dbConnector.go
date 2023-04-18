package common

import (
	"context"
	"database/sql"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/exp/slog"
)

type DbConnector struct {
	AutoCloseable
	db  *bun.DB
	ctx *context.Context
}

var dbCon *DbConnector
var dbOnce sync.Once

func GetDbCon() *DbConnector {
	dbOnce.Do(func() {
		dbCon = &DbConnector{}
	})
	return dbCon
}

func (dbCon *DbConnector) Connect(dsn string, ctx *context.Context) error {
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	dbCon.db = bun.NewDB(sqlDb, pgdialect.New())
	dbCon.ctx = ctx
	slog.Info("connect to:{} success", dsn)
	Get().Add(dbCon)
	return nil
}

func (dbCon *DbConnector) GetDb() (*bun.DB, *context.Context) {
	return dbCon.db, dbCon.ctx
}
