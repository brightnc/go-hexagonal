package repository

import (
	"database/sql"

	"github.com/brightnc/go-hexagonal/internal/core/domain"
	"github.com/brightnc/go-hexagonal/internal/core/port"
	"github.com/jmoiron/sqlx"
)

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) port.AccountRepository {
	return &accountRepositoryDB{
		db: db,
	}
}

func (r *accountRepositoryDB) Create(acc domain.Account) (*domain.Account, error) {

	query := "insert into accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)

	return &acc, nil
}

func (r *accountRepositoryDB) GetAll(customerID int) ([]domain.Account, error) {
	var accounts []domain.Account
	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id=?"
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, sql.ErrNoRows
	}
	return accounts, nil
}
