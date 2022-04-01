package database

import (
	"database/sql"
	"fmt"

	// importa o driver e usa o _ para deixar implicita
	_ "github.com/go-sql-driver/mysql"
)

type dbConnection struct {
	user   string
	pass   string
	host   string
	port   int
	dbName string
}

var db = dbConnection{
	user:   "root",
	pass:   "wykx@5jd9",
	host:   "localhost",
	port:   33306,
	dbName: "golangmysql",
}

func Connect() (*sql.DB, error) {
	mysqlUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", db.user, db.pass, db.host, db.port, db.dbName)
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
