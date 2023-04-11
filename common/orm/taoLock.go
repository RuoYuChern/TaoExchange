package orm

import (
	"time"

	"github.com/uptrace/bun"
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

//func (tlk *TaoLock)insert()
