package postgresql

import (
	"VkProject/internal/config"
	"VkProject/pkg/utils"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	//BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	//BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(tx pgx.Tx) error) error
}

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	err = repeatable.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {

		log.Fatal("error do with tries postgresql")
	}

	return pool, nil
}
