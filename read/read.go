package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type DepositCalculation struct {
	BankName    string  `json:"bank_name,omitempty"`
	Account     string  `json:"account,omitempty"`
	AccountType string  `json:"account_type,omitempty"`
	Apy         float64 `json:"apy,omitempty"`
	Years       float64 `json:"years,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Delta       float64 `json:"delta,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need a Kafka topic name.")
		return
	}

	partition := 0
	topic := os.Args[1]
	fmt.Println("Kafka topic:", topic)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	defer r.Close()
	err := r.SetOffset(0)
	if err != nil {
		log.Fatalf("error setting offset %v", err)
	}

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("err when reading message is %v", err)
			continue
		}
		fmt.Printf("message at offset %d: MKey %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		temp := DepositCalculation{}
		err = json.Unmarshal(m.Value, &temp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%T\n", temp)
	}
}
