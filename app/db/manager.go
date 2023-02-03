package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Manager struct {
	db *sql.DB
}

func newDbConnection(username, password, host, DBName string, port int) *sql.DB {
	// Open a connection to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, DBName))
	if err != nil {
		log.Println(err)
		return db
	}
	return db
}

func NewManager(username, password, host, DBName string, port int) *Manager {
	return &Manager{
		db: newDbConnection(username, password, host, DBName, port),
	}
}
