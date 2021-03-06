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
	topic := "depositcalculation-grpc"
	fmt.Println("IlluminatingDeposits gRPC Kafka topic:", topic)
	partition := 0

	kafkaURL := "kafka:9092"
	if khost, ok := os.LookupEnv("DEPOSITS_GRPC_KAFKA_HOST"); ok {
		kafkaURL = fmt.Sprintf("%v:9092", khost)
	}
	log.Println("kafkaURL is (should match what is used for producer)", kafkaURL)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaURL},
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
