package repository

import (
	"fmt"

	"github.com/brightnc/go-hexagonal/internal/core/domain"
	"github.com/brightnc/go-hexagonal/internal/core/port"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) port.CustomerRepository {
	return &customerRepositoryDB{
		db: db,
	}
}

func (r *customerRepositoryDB) GetAll() ([]domain.Customer, error) {
	var users []domain.Customer
	query := "SELECT id, user_name, password , email, created_at, updated_at FROM users"
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}
func (r *customerRepositoryDB) GetById(id int) (*domain.Customer, error) {
	var user domain.Customer
	query := "SELECT id, user_name, password , email, status, created_at, updated_at FROM users WHERE id = ?"
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *customerRepositoryDB) CreateUser(user domain.Customer) (*domain.Customer, error) {
	query := "insert into users (id, user_name, password, email,status, created_at, updated_at) values (?, ?, ?, ?, ?, ?,?)"

	// hashed password

	hasedByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user.Password = string(hasedByte)

	result, err := r.db.Exec(query, user.ID, user.Username, user.Password, user.Email, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return &user, nil
}
