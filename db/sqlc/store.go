package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and its composite transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err:%v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	SourceAccountID int64 `json:"SourceAccountID"`
	TargetAccountID int64 `json:"TargetAccountID"`
	Amount          int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer      Transfer `json:"transfer"`
	SourceAccount Account  `json:"SourceAccount"`
	TargetAccount Account  `json:"TargetAccount"`
	FromEntry     Entry    `json:"FromEntry"`
	ToEntry       Entry    `json:"ToEntry"`
}

var txKey = struct{}{}

// TransferTx performs a money transfer from one account to another
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, "create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			SourceAccountID: arg.SourceAccountID,
			TargetAccountID: arg.TargetAccountID,
			Amount:          arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.SourceAccountID,
			Amount:    -1 * arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.TargetAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.SourceAccountID < arg.TargetAccountID {
			result.SourceAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.SourceAccountID,
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}

			result.TargetAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.TargetAccountID,
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}
		} else {
			result.TargetAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.TargetAccountID,
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}

			result.SourceAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.SourceAccountID,
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}
		}

		return nil

	})

	return result, err

}
