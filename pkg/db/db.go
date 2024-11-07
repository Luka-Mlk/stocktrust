package db

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

var connecton *pgx.Conn

func Conn() (*pgx.Conn, error) {
	if connecton == nil {
		user := os.Getenv("DATABASE_USER")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABAES_PORT")
		dbname := os.Getenv("DATABASE_NAME")
		connstr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
		pgxCon, err := pgx.ParseConnectionString(connstr)
		if err != nil {
			return nil, err
		}
		conn, err := pgx.Connect(pgxCon)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		connecton = conn
	}
	return connecton, nil
}
