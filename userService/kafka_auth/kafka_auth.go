package kafka_auth

import (
	"log"
	"userService/api"
	"userService/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func CheckIFAuthorized(server *api.Server) {
	var is_authorized string
	req_topic := "check_auth"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "auth_group_1",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		log.Fatal(err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		log.Fatal(err)
	}

	res_topic := "is_authorized"
	delivery_chan := make(chan kafka.Event, 100)

	err = consumer.Subscribe(req_topic, nil)
	is_authorized = "false"
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {

		case *kafka.Message:
			if string(e.Value) != "" {

				//-------------CHECK IF AUTHORIZED-----------//
				log.Println(string(e.Value))

				claims, err := utils.ValidateToken(string(e.Value), server.Config.JWT_SECRET)

				if err != "" {
					log.Fatal(err)
				}
				log.Println(claims.Email)
				if claims.Email != "" {
					is_authorized = "true"
				}
				//-------------------------------------------//

				err2 := producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{
						Topic:     &res_topic,
						Partition: kafka.PartitionAny,
					},
					Value: []byte(is_authorized)},
					delivery_chan,
				)
				if err2 != nil {
					log.Fatal(err)
				}

			}

			<-delivery_chan

		case *kafka.Error:
			log.Fatal(err)
		}
	}

}
