package database

import (
	"database/sql"
	"fmt"

	"github.com/go-postgresql-crud/config"
	"github.com/go-postgresql-crud/logger"
	"github.com/jmoiron/sqlx"

	//postgresql driver
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// DbObject
var DbObject *sqlx.DB

//Pinger ...
type Pinger interface {
	Ping() error
}

//Connect ...
func Connect() {

	dbconnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.ConfigKeys.DbHost,
		config.ConfigKeys.DbPort,
		config.ConfigKeys.DbUser,
		config.ConfigKeys.DbPassword,
		config.ConfigKeys.DbName)

	db, err := sqlx.Connect("postgres", dbconnection)
	if err != nil {
		logger.Error(err, "Connect to database")
	} else {
		DbObject = db.Unsafe()
	}
}

//DB ...
type DB struct {
	Pinger Pinger
}

// Ping : Ping the db
func (db DB) Ping() error {

	err := db.Pinger.Ping()
	if err != nil {
		err = errors.Wrap(err, "Ping")
		logger.Error(err, "")
	}
	return err
}

// Close ...
func Close() {
	DbObject.Close()
}

//Database interface
type Database interface {
	Select(dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
}
