package orm

import (
	"tao.exchange.com/infra"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/exp/slog"
)

const (
	FREE_ST      = 0
	LOCK_ST      = 1
	UN_HEALTH_ST = 2
)

type TaoLock struct {
	bun.BaseModel `bun:"table:tao_lock"`
	Id            int32     `bun:"id,autoincrement"`
	ShardId       string    `bun:"shardid,pk"`
	AppId         string    `bun:"appid,notnull"`
	AppIP         string    `bun:"appip,notnull"`
	AppRole       string    `bun:"approle,notnull"`
	AppPort       int32     `bun:"appport,notnull"`
	AppStatus     int32     `bun:"appstatus,notnull"`
	LockTime      time.Time `bun:"locktime,notnull,default:current_timestamp"`
}

type TaoLockMapper struct {
}

func (tlm *TaoLockMapper) BatchSelect() ([]TaoLock, error) {
	db, ctx := infra.GetDbCon().GetDb()
	taoLockList := make([]TaoLock, 0)
	err := db.NewRaw("SELECT id,shardid,appid,appip,approle,appport,appstatus FROM ? WHERE appstatus = ?",
		bun.Ident("tao_lock"), LOCK_ST).Scan(*ctx, &taoLockList)
	if err != nil {
		slog.Info("BatchSelect:", err.Error())
		return nil, err
	}
	return taoLockList, nil

}

func (tlm *TaoLockMapper) ReleaseLock(tlk *TaoLock) int32 {
	db, ctx := infra.GetDbCon().GetDb()
	r, err := db.NewUpdate().Model(tlk).Column("appstatus", "locktime").Where("shardid = ?", tlk.ShardId).Where("appstatus = ?", LOCK_ST).
		Where("appid = ?", tlk.AppId).Exec(*ctx)
	if err != nil {
		slog.Info("ReleaseLock error:", err.Error())
		return -1
	}

	v, err := r.RowsAffected()
	if err != nil {
		slog.Info("ReleaseLock error:", err.Error())
		return -1
	}
	return int32(v)
}

func (tlm *TaoLockMapper) Insert(tlk *TaoLock) int32 {
	db, ctx := infra.GetDbCon().GetDb()
	r, err := db.NewInsert().Model(tlk).On("CONFLICT (shardId) DO UPDATE").Set("appid = EXCLUDED.appid").
		Set("appip = EXCLUDED.appip").Set("approle = EXCLUDED.approle").
		Set("appport = EXCLUDED.appport").Set("appstatus = EXCLUDED.appstatus").Set("locktime = EXCLUDED.locktime").
		Where("appStatus == 0").Exec(*ctx)
	if err != nil {
		slog.Info("Insert error:", err.Error())
		return -1
	}
	v, err := r.RowsAffected()
	if err != nil {
		slog.Info("ReleaseLock error:", err.Error())
		return -1
	}
	return int32(v)

}
