package main

import (
	"database/sql"
	"tax-service/internal/order/infra/database"
	"tax-service/internal/order/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 100,
		Tax:   10,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}

	println(output.FinalPrice)
}
