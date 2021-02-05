package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

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
	TOTAL := 15
	topic := ""
	if len(os.Args) > 0 {
		topic = os.Args[1]
	} else {
		fmt.Println("Usage:", os.Args[0], "TOPIC")
		return
	}

	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	defer conn.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	rand.Seed(time.Now().Unix())

	for i := 0; i < TOTAL; i++ {
		temp := DepositCalculation{
			BankName:    "LINKED",
			Account:     fmt.Sprint("2552", i),
			AccountType: "CD",
			Apy:         2.22,
			Years:       8,
			Amount:      12552,
		}
		recordJSON, _ := json.Marshal(temp)

		err := conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			log.Fatal("failed to set write deadline:", err)
		}
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(recordJSON)},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if i%5 == 0 {
			fmt.Print("5DepositsSent.")
		}
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println()
}
