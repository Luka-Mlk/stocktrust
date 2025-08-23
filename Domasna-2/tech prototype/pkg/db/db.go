package db

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
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
		connstr := fmt.Sprintf("your connection string here")
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
			log.Println(err)
			debug.PrintStack()
			return nil, err
		}
	}
	conn, err := conPool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return nil, err
	}
	return conn, nil
}
