package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"tax-service/internal/order/infra/database"
	"tax-service/internal/order/usecase"
	"tax-service/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	maxWorkers := 3

	wg := sync.WaitGroup{}

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewGetTotalUseCase(repository)
		output, err := uc.Execute()
		if err != nil {
			// Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(output)
	})
	go http.ListenAndServe(":8181", nil)

	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)

	wg.Add(maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go worker(out, uc, i)
	}

	wg.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {

		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}

		input.Tax = 10.0

		_, err = uc.Execute(input)

		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}

		msg.Ack(false)

		fmt.Println("Worker", workerId, "processed order", input.ID)
	}
}
