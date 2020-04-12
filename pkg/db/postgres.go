package db

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

import (
	_ "github.com/jackc/pgx/stdlib"
)

func PostgresOpenConnection(user, pass, host, db string) (*sqlx.DB, error) {
	conn, errOpen := sqlx.Open(`postgres`, fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True",
		user,
		pass,
		host,
		db,
	))

	if errOpen != nil {
		return nil, errOpen
	}
	if errPing := conn.Ping(); errPing != nil {
		return nil, errPing
	}

	return conn, nil
}
