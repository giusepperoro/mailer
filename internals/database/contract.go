package database

import "context"

type DbManager interface {
	ChangeBalance(ctx context.Context, clientId int64, amount int64) (bool, error)
}
