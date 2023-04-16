package kafka_auth

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

func CheckIFAuthorized(ctx *gin.Context) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "auth_group",
		"auto.offset.reset": "smallest",
	})
	consumer_topic := "is_authorized"

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	token := ctx.GetHeader("Authorization")

	if token == "" {

		log.Fatal("token not found")
	}
	log.Println(token)

	producer_topic := "check_auth"
	delivery_chan := make(chan kafka.Event, 100)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &producer_topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(token)},
		delivery_chan,
	)

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	<-delivery_chan

	err = consumer.Subscribe(consumer_topic, nil)
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {

		case *kafka.Message:

			log.Println(string(e.Value))

		case *kafka.Error:
			log.Fatal(err)
		}
	}

}
