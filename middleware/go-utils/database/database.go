// Package database provides gorm connection
package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql" // initate go sql driver
)

var (
	// DBConn Current Database connection
	DBConn       *gorm.DB
	ClientInfoDb *gorm.DB
	// Err Database connection error
	Err error
)

// MySQLConnect Connect to a MySQL driver-based database
func MySQLConnect(username, password, host, databaseName string) {
	if host != "" {
		host = fmt.Sprintf("tcp(%s)", host)
	}

	DBConn, Err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@%s/%s?parseTime=true", username, password, host, databaseName)), &gorm.Config{})
}

// SQLiteConnect ...
func SQLiteConnect(filename string) {
	DBConn, Err = gorm.Open(sqlite.Open(filename), &gorm.Config{})
}

// PostgreSQLConnect Connect to a PostgreSQL database
func PostgreSQLConnect(username, password, host, databaseName, port, sslMode, timeZone string) {
	DBConn, Err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, username, password, databaseName, port, sslMode, timeZone)), &gorm.Config{})
}

type Database struct {
	ClientInfoDb *gorm.DB
}

var Data Database

func ConnectDB() {
	ClientInfoDns := "host=35.241.103.247 user=fdsap-ellaine password=fdsap@2024 dbname=FDS-CORE port=18010 sslmode=disable"
	ClientInfoDb, err := gorm.Open(postgres.Open(ClientInfoDns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect in database \n", err.Error())
		os.Exit(2)
	}

	// err = ClientInfoDb.Exec("SET search_path TO ewallet").Error
	// if err != nil {
	//     log.Fatal("Failed to set search path\n", err.Error())
	//     os.Exit(2)
	// }

	// ClientInfoDb.Logger = logger.Default.LogMode(logger.Info)
	// log.Println("Running Migration for database ClientInfo")
	// // if err := ClientInfoDb.AutoMigrate(); err != nil {
	// // 	log.Fatalf("Error running migration for database ClientInfo %v", err)
	// // }

	Data = Database{
		ClientInfoDb: ClientInfoDb,
	}
}
