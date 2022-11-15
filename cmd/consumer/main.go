package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ab-costa/imersao-go/internal/order/infra/database"
	"github.com/ab-costa/imersao-go/internal/order/pkg/rabbitmq"
	"github.com/ab-costa/imersao-go/internal/order/usecase"
	amqp "github.com/rabbitmq/amqp091-go"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// T1
func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUsecase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	out := make(chan amqp.Delivery) // CHANNEL

	// T2
	go rabbitmq.Consumer(ch, out)

	// T1
	for msg := range out {
		var inputDTO usecase.OrderInputDTO

		err := json.Unmarshal(msg.Body, &inputDTO)

		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)

		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println(outputDTO)
	}
}