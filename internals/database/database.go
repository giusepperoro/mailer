package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type dataBase struct {
	conn *pgxpool.Pool
}

func New(ctx context.Context) (*dataBase, error) {

	connection, err := pgxpool.Connect(ctx, "postgres://postgres:postgres@0.0.0.0:5432/master")
	if err != nil {
		return nil, err
	}
	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") == "1" {
		batch := &pgx.Batch{}
		queryStr := `INSERT INTO accounts (balance) VALUES ($1) `
		n := rand.Intn(3000)
		for i := 0; i < n; i++ {
			randAmount := rand.Intn(10000) + 10000
			batch.Queue(queryStr, randAmount)
		}
		c, err := connection.Acquire(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to create clients for debug: %w", err)
		}
		br := c.SendBatch(ctx, batch)
		err = br.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to create clients for debug: %w", err)
		}
		log.Println("test clients created...")
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
