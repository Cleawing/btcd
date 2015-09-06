package cassandra

import (
	"github.com/btcsuite/btcd/database"
)

func init() {
	driver := database.DriverDB{DbType: "cassandra", CreateDB: CreateDB, OpenDB: OpenDB}
	database.AddDBDriver(driver)
}

func OpenDB(args ...interface{}) (database.Db, error) {
	return nil, nil
}

func CreateDB(args ...interface{}) (database.Db, error) {
	db := Cassandra{}
	return &db, nil
}