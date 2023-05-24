package database

import (
	"database/sql"
	"fmt"
	"log"
	"simple-gin-server/config"
)

const DATABASE = "mysql"

var databaseSession *sql.DB

func getDatabaseDSN() string {
	conf := config.GetConfig()
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)
}

func Init() {
	dsn := getDatabaseDSN()
	db, err := sql.Open(DATABASE, dsn)
	if err != nil {
		log.Fatalf("Error invoked with MySQL session: %s", err.Error())
	}
	databaseSession = db
}

func GetDatabaseSession() *sql.DB {
	return databaseSession
}
