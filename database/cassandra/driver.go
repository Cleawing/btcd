package cassandra

import (
	"github.com/btcsuite/btcd/database"
	"github.com/hailocab/gocassa"
	"github.com/btcsuite/btclog"
)

var log = btclog.Disabled

func init() {
	driver := database.DriverDB{DbType: "cassandra", CreateDB: CreateDB, OpenDB: OpenDB}
	database.AddDBDriver(driver)
}

func OpenDB(args ...interface{}) (database.Db, error) {
	log = database.GetLog()
	var tables []string

	connection, err := gocassa.Connect([]string{"192.168.99.101"}, "cassandra", "")

	if err != nil {
		return nil, err
	}

	keySpace := connection.KeySpace("bitcoin")
	tables, err = keySpace.Tables()

	if err != nil {
		return nil, err
	}

	if len(tables) == 0 {
		log.Warn("Empty keyspace")
	}

	db := Cassandra{
		keySpace: keySpace,
	}

	return &db, nil
}

func CreateDB(args ...interface{}) (database.Db, error) {
	keySpace, err := gocassa.ConnectToKeySpace("bitcoin", []string{"192.168.99.101"}, "cassandra", "")

	if err != nil {
		return nil, err
	}
	db := Cassandra{
		keySpace: keySpace,
	}
	return &db, nil
}