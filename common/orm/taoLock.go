package orm

import (
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
)

const (
	FREE_ST      = 0
	LOCK_ST      = 1
	UN_HEALTH_ST = 2
)

type TaoLock struct {
	bun.BaseModel `bun:"table:tao_lock"`
	Id            int32     `bun:"id,autoincrement"`
	ShardId       string    `bun:"shardId,pk"`
	AppId         string    `bun:"appId,notnull"`
	AppIP         string    `bun:"appIP,notnull"`
	AppRole       string    `bun:"appRole,notnull"`
	AppPort       int32     `bun:"appPort,notnull"`
	AppStatus     int32     `bun:"appStatus,notnull"`
	LockTime      time.Time `bun:"lockTime,notnull,default:current_timestamp"`
}

type TaoLockMapper struct {
}

func (tlm *TaoLockMapper) Insert(tlk *TaoLock) int32 {
	db, ctx := common.GetDbCon().GetDb()
	_, err := db.NewInsert().Model(tlk).Column("shardId", "appId", "appIp", "appRole", "appPort", "appStatus", "lockTime").
		On("CONFLICT (shardId) DO UPDATE").Set("appId = EXCLUDED.appId").Where("appStatus == 0").Exec(*ctx)
	if err != nil {
		slog.Info("")
		return -1
	}

	return 1
}
