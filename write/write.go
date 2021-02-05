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

type Record struct {
	Name   string `json:"name"`
	// Random int    `json:"random"`
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	// MIN := 0
	// MAX := 0
	TOTAL := 500
	topic := ""
	if len(os.Args) > 0 {
		// MIN, _ = strconv.Atoi(os.Args[1])
		// MAX, _ = strconv.Atoi(os.Args[2])
		// TOTAL, _ = strconv.Atoi(os.Args[1])
		topic = os.Args[1]
	} else {
		fmt.Println("Usage:", os.Args[0], "TOTAL TOPIC")
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
		// myrand := random(MIN, MAX)
		// temp := Record{strconv.Itoa(i), myrand}
		temp := Record{Name: fmt.Sprintf("rohit%v", i)}
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
