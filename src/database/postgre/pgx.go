package postgre

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafitanujaya/go-fiber-template/src/config"
	"github.com/samber/do/v2"
)

func NewPgxConnectInject(i do.Injector) (*pgxpool.Pool, error) {
	DbString := config.GetDBConnection()
	fmt.Printf("Connection string: %s\n", DbString)
	tempConn, err := pgx.Connect(context.Background(), DbString)
	if err != nil {
		panic(err.Error())
	}
	defer tempConn.Close(context.Background())

	var maxConStr string

	row := tempConn.QueryRow(context.Background(), "SHOW max_connections")
	err = row.Scan(&maxConStr)

	if err != nil {
		panic(err)
	}

	maxConn, err := strconv.Atoi(maxConStr)
	if err != nil {
		panic(err)
	}

	maxOpenConnection := int(float64(maxConn) * 0.8)

	pgxConfig, err := pgxpool.ParseConfig(DbString)
	if err != nil {
		panic(err)
	}

	pgxConfig.MaxConns = int32(maxOpenConnection)

	db, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		panic(err)
	}

	return db, nil
}
