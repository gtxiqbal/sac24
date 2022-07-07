package config

import (
	"database/sql"
	"fmt"
	"github.com/gtxiqbal/sac24/helper"
	_ "github.com/lib/pq"
	"time"
)

func NewPG(dbDriver, dbHost, dbName, dbUsername, dbPassword string, dbPort int) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbPassword, dbName)
	db, err := sql.Open(dbDriver, psqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
