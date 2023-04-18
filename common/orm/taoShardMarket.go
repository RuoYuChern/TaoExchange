package orm

import (
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
)

type TaoShardMarket struct {
	bun.BaseModel `bun:"table:tao_shard_market"`
	Id            int32     `bun:"id,autoincrement"`
	MarketId      string    `bun:"marketid,pk"`
	ShardId       string    `bun:"shardid,notnull"`
	CreateTime    time.Time `bun:"createtime,notnull,default:current_timestamp"`
}

type TaoShardMarketMapper struct {
}

func (tsmm *TaoShardMarketMapper) BatchSelect(id int32) ([]TaoShardMarket, error) {
	tsms := make([]TaoShardMarket, 0)
	db, ctx := common.GetDbCon().GetDb()
	err := db.NewRaw("SELECT id,marketid,shardid,createtime FROM ? WHERE id >? ORDER BY id asc LIMIT ?", bun.Ident("tao_shard_market"), id, 1000).Scan(*ctx, &tsms)
	if err != nil {
		slog.Info("BatchSelect:", err.Error())
		return nil, err
	}
	return tsms, nil
}

func (tsmm *TaoShardMarketMapper) Insert(tsm *TaoShardMarket) int32 {
	db, ctx := common.GetDbCon().GetDb()
	_, err := db.NewInsert().Model(tsm).On("CONFLICT (marketid) DO UPDATE").Set("shardid = EXCLUDED.shardid").Set("createtime = EXCLUDED.createtime").Exec(*ctx)
	if err != nil {
		slog.Info("Insert:", err.Error())
		return -1
	}

	return 1
}

func (tsmm *TaoShardMarketMapper) BachInsert(tsmList *[]TaoShardMarket) int32 {
	db, ctx := common.GetDbCon().GetDb()
	_, err := db.NewInsert().Model(tsmList).On("CONFLICT (marketid) DO UPDATE").Set("shardid = EXCLUDED.shardid").Set("createtime = EXCLUDED.createtime").Exec(*ctx)
	if err != nil {
		slog.Info("BachInsert:", err.Error())
		return -1
	}
	return int32(len(*tsmList))
}
