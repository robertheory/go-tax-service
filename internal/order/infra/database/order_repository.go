package database

import (
	"database/sql"
	"tax-service/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (repo *OrderRepository) Save(order *entity.Order) error {

	_, err := repo.Db.Prepare("INSERT INTO orders (id, price, tax) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = repo.Db.Exec(order.ID, order.Price, order.Tax)

	if err != nil {
		return err
	}

	return nil
}
