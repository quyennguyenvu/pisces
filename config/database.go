package config

import (
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"           // mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgresql driver
)

// DBConnection ..
type DBConnection struct {
	Driver     string
	DataSource string
	Database   string
	LogMode    bool
}

var confDB *DBConnection
var dbInit sync.Once

// GetConnection ..
func GetConnection() *DBConnection {
	dbInit.Do(func() {
		drive := (*confEnv)["DB_CONNECTION"]
		dbHost := (*confEnv)["DB_HOST"]
		dbPort := (*confEnv)["DB_PORT"]
		database := (*confEnv)["DB_DATABASE"]
		userName := (*confEnv)["DB_USERNAME"]
		password := (*confEnv)["DB_PASSWORD"]

		// select datasource
		var dataSource string
		switch drive {
		case "mysql":
			dataSource = userName + ":" + password +
				"@tcp(" + dbHost + ":" + dbPort + ")/" + database +
				"?charset=utf8&parseTime=True&loc=Local"
		case "postgres":
			dataSource = "host=" + dbHost +
				" port=" + dbPort +
				" user=" + userName +
				" dbname=" + database +
				" password=" + password +
				" sslmode=disable"
		}

		logMode, _ := strconv.ParseBool((*confEnv)["DB_LOGMODE"])

		confDB = &DBConnection{
			Driver:     drive,
			DataSource: dataSource,
			Database:   database,
			LogMode:    logMode,
		}
	})
	return confDB
}
