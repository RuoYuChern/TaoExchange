package common

type AutoCloseable interface {
	AutoClose()
}

func (dbCon *DbConnector) AutoClose() {
	if dbCon.db != nil {
		dbCon.db.Close()
	}
}
