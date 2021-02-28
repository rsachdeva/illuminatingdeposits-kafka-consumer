package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

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
	writeHighLevelAPI()
}

func writeHighLevelAPI() {
	topic := "depositcalculationhigh-1"
	kafkaURL := "127.0.0.1:9092"
	createTopic(kafkaURL, topic)
	kafkaWriter := getKafkaWriter(kafkaURL, topic)
	log.Println("kafkaWriter is ", kafkaWriter)
	defer func() {
		err := kafkaWriter.Close()
		if err != nil {
			log.Fatalln("could not close depositcalculation connection", err)
		}
	}()

	svc := ServiceServer{
		KafkaWriter: kafkaWriter,
	}

	dc := DepositCalculation{
		BankName:    "LINKED",
		Account:     fmt.Sprint("2552", 1000),
		AccountType: "CD",
		Apy:         2.22,
		Years:       8,
		Amount:      12552,
	}
	svc.writeMessage(context.Background(), dc)
}

func createTopic(kafkaURL, topic string) {
	conn, err := kafka.Dial("tcp", kafkaURL)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	log.Printf("kafka.TCP(kafkaURL) is %v", kafka.TCP(kafkaURL))
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

type ServiceServer struct {
	KafkaWriter *kafka.Writer
}

func (svc ServiceServer) writeMessage(ctx context.Context, dc DepositCalculation) {
	recordJSON, _ := json.Marshal(dc)
	log.Println("svc.AccessToken is ", "jmd")
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("accessToken-%s", "jmdshivganesha2")),
		Value: recordJSON,
	}
	err := svc.KafkaWriter.WriteMessages(ctx, msg)

	if err != nil {
		log.Fatalln("could not write to kafka writer:", err)
	}
}

// func writeLowLevelAPI() {
// 	TOTAL := 15
// 	topic := "depostlowlevel"
// 	// if len(os.Args) > 0 {
// 	// 	topic = os.Args[1]
// 	// } else {
// 	// 	fmt.Println("Usage:", os.Args[0], "TOPIC")
// 	// 	return
// 	// }
//
// 	partition := 0
// 	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
// 	defer conn.Close()
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 		return
// 	}
//
// 	rand.Seed(time.Now().Unix())
//
// 	for i := 0; i < TOTAL; i++ {
// 		temp := DepositCalculation{
// 			BankName:    "LINKED",
// 			Account:     fmt.Sprint("2552", i),
// 			AccountType: "CD",
// 			Apy:         2.22,
// 			Years:       8,
// 			Amount:      12552,
// 		}
// 		recordJSON, _ := json.Marshal(temp)
//
// 		err := conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
// 		if err != nil {
// 			log.Fatal("failed to set write deadline:", err)
// 		}
// 		_, err = conn.WriteMessages(
// 			kafka.Message{Value: []byte(recordJSON)},
// 		)
// 		if err != nil {
// 			log.Fatal("failed to write messages:", err)
// 		}
//
// 		if i%5 == 0 {
// 			fmt.Print("5DepositsSent.")
// 		}
// 		time.Sleep(10 * time.Millisecond)
// 	}
//
// 	fmt.Println()
// }
