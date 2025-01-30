package main

import (
	"database/sql"
	"fmt"
	"tax-service/internal/order/entity"
	"tax-service/internal/order/infra/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	order, err := entity.NewOrder("2", 20, 2)

	if err != nil {
		panic(err)
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		panic(err)
	}

	fmt.Printf("The final price is: %f", order.FinalPrice)

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	repository := database.NewOrderRepository(db)

	err = repository.Save(order)

	if err != nil {
		panic(err)
	}

}
