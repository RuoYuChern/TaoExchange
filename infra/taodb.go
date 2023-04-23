package infra

import (
	"context"
	"database/sql"
	"sync"
	"tao.exchange.com/common"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/exp/slog"
)

type TaoDb struct {
	common.TaoCloseable
	db  *bun.DB
	ctx *context.Context
}

var dbCon *TaoDb
var dbOnce sync.Once

func GetDbCon() *TaoDb {
	dbOnce.Do(func() {
		dbCon = &TaoDb{}
	})
	return dbCon
}

func (dbCon *TaoDb) Connect(dsn string, ctx *context.Context) error {
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	dbCon.db = bun.NewDB(sqlDb, pgdialect.New())
	dbCon.ctx = ctx
	slog.Info("connect to:{} success", dsn)
	common.Get().Add(dbCon)
	return nil
}

func (dbCon *TaoDb) GetDb() (*bun.DB, *context.Context) {
	return dbCon.db, dbCon.ctx
}

func (dbCon *TaoDb) AutoClose() {
	if dbCon.db != nil {
		slog.Info("DbConnector close")
		err := dbCon.db.Close()
		if err != nil {
			return
		}
	}
}
