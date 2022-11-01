package database

import (
	"context"
	pgx "github.com/jackc/pgx/v4"
)

type dataBase struct {
	conn *pgx.Conn
}

func New(ctx context.Context) (*dataBase, error) {

	connection, err := pgx.Connect(ctx, "postgres://postgres:postgres@database:5432/master")
	if err != nil {
		return nil, err
	}
	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &dataBase{conn: connection}, nil
}
