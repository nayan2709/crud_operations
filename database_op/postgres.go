package database_op

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Database struct {
	host     string
	port     int
	user     string
	password string
	database string
}

func NewDatabase(host string, port int, user string, password string, database string) Database {
	return Database{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
	}
}

func (db *Database) InitializePostgres() (*gorm.DB, error) {
	pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.host, db.port, db.user, db.password, db.database)
	database, err := gorm.Open("postgres", pgConnStr)
	if err != nil {
		return nil, err
	}
	database.LogMode(true)
	return database, nil
}
