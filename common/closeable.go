package common

import "golang.org/x/exp/slog"

type AutoCloseable interface {
	AutoClose()
}

func (dbCon *DbConnector) AutoClose() {
	if dbCon.db != nil {
		slog.Info("DbConnector close")
		dbCon.db.Close()
	}
}
