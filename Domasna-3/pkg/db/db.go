package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var conPool *pgxpool.Pool

func Conn() (*pgxpool.Conn, error) {
	const defaultMaxConns = int32(40)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	if conPool == nil {
		user := os.Getenv("DATABASE_USER")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABAES_PORT")
		dbname := os.Getenv("DATABASE_NAME")
		connstr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
		pgxConf, err := pgxpool.ParseConfig(connstr)
		if err != nil {
			return nil, err
		}
		pgxConf.MaxConns = defaultMaxConns
		pgxConf.MinConns = defaultMinConns
		pgxConf.MaxConnLifetime = defaultMaxConnLifetime
		pgxConf.MaxConnIdleTime = defaultMaxConnIdleTime
		pgxConf.HealthCheckPeriod = defaultHealthCheckPeriod
		pgxConf.ConnConfig.ConnectTimeout = defaultConnectTimeout
		conPool, err = pgxpool.NewWithConfig(context.Background(), pgxConf)
		if err != nil {
			e := fmt.Errorf("error creating new pool:\n%v", err)
			return nil, e
		}
	}
	conn, err := conPool.Acquire(context.Background())
	if err != nil {
		e := fmt.Errorf("error aquire connection from pool:\n%v", err)
		return nil, e
	}
	return conn, nil
}
