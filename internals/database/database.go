package database

import (
	"context"
	"errors"
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

func (d *dataBase) ChangeBalance(ctx context.Context, clientId int64, amount int64) (bool, error) {
	var balance, id int64
	var query = "SELECT balance, client_id FROM accounts WHERE client_id = $1"
	row := d.conn.QueryRow(ctx, query, clientId)
	err := row.Scan(&balance, &id)
	if err != nil {
		return false, err
	}
	if id != clientId {
		return false, errors.New("client does not exist")
	}

	if balance+amount < 0 {
		return false, errors.New("negative balance")
	}

	query = "UPDATE accounts SET balance = balance + $1 WHERE client_id = $2"
	d.conn.QueryRow(ctx, query, amount, clientId)
	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}
	return true, nil
}
