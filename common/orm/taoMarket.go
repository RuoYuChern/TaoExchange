package orm

import (
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/exp/slog"
	"tao.exchange.com/common"
)

type TaoMarket struct {
	bun.BaseModel `bun:"table:tao_market"`
	Id            int32     `bun:"id,autoincrement"`
	Market        string    `bun:"market,pk"`
	Base          string    `bun:"base,notnull"`
	Pair          string    `bun:"pair,notnull"`
	MarketStatus  int32     `bun:"marketstatus,notnull"`
	BasePrec      int32     `bun:"baseprec,notnull"`
	PairPrec      int32     `bun:"pairprec,notnull"`
	FeePrec       int32     `bun:"feeprec,notnull"`
	MinAmount     int64     `bun:"minamount,notnull"`
	MinBase       int64     `bun:"minbase,notnull"`
	CreateTime    time.Time `bun:"createtime,notnull,default:current_timestamp"`
}

type TaoMarketMapper struct {
}

func (tmm *TaoMarketMapper) BachInsert(tms []TaoMarket) (int32, error) {
	db, ctx := common.GetDbCon().GetDb()
	_, err := db.NewInsert().Model(tms).On("CONFLICT (market) DO UPDATE").Set("base = EXCLUDED.base").
		Set("pair = EXCLUDED.pair").Set("marketstatus = EXCLUDED.marketstatus").
		Set("baseprec = EXCLUDED.baseprec").Set("pairprec = EXCLUDED.pairprec").Set("feeprec = EXCLUDED.feeprec").
		Set("minamount = EXCLUDED.minamount").Set("minbase = EXCLUDED.minbase").Set("createtime = EXCLUDED.createtime").Exec(*ctx)
	if err != nil {
		slog.Info("BachInsert:", err.Error())
		return -1, err
	}
	return int32(len(tms)), nil
}

func (tmm *TaoMarketMapper) BatchSelect(id int32) ([]TaoMarket, error) {
	db, ctx := common.GetDbCon().GetDb()
	tmks := make([]TaoMarket, 0)
	err := db.NewRaw("SELECT id,market,base,pair,marketstatus,baseprec,pairprec,feeprec,minamount,minbase,createtime FROM ? WHERE id >? ORDER BY id asc LIMIT ?",
		bun.Ident("tao_market"), id, 1000).Scan(*ctx, &tmks)
	if err != nil {
		slog.Info("BatchSelect:", err.Error())
		return nil, err
	}
	return tmks, nil
}
