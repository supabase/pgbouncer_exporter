package main

import (
	"database/sql"
	"fmt"
)

type DbConn struct {
	db         *sql.DB
	connString string
}

func (d *DbConn) GetDb() (*sql.DB, error) {
	if d.db != nil {
		return d.db, nil
	}
	db, err := sql.Open("postgres", d.connString)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SHOW STATS")
	if err != nil {
		return nil, fmt.Errorf("error pinging pgbouncer: %q", err)
	}
	defer rows.Close()

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	return db, nil
}
