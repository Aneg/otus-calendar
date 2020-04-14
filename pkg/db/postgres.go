package db

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func PostgresOpenConnection(user, pass, host, port, db string) (*sqlx.DB, error) {
	conn, errOpen := sqlx.Open(
		`pgx`,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			user,
			pass,
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
