package config

import (
	"database/sql"
	"fmt"
	"github.com/gtxiqbal/sac24/helper"
	_ "github.com/lib/pq"
	"time"
)

func NewPG(dbDriver, dbHost, dbName, dbUsername, dbPassword string,
	dbPort, dbMaxIdleConn, dbMaxOpenConn, dbConnMaxIdleTime, dbConnMaxLifeTime int) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbPassword, dbName)
	db, err := sql.Open(dbDriver, psqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(dbMaxIdleConn)
	db.SetMaxOpenConns(dbMaxOpenConn)
	db.SetConnMaxIdleTime(time.Duration(dbConnMaxIdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(dbConnMaxLifeTime) * time.Minute)

	return db
}
