package kafkaconsumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OTP_details struct {
	Otp   string
	Email string
}

func ConsumeOTP() {
	topic := "OTP"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "otp_group",
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Subscribe(topic, nil)

	if err != nil {
		log.Fatal(err)
	}

	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			//decode the byte message received from kafka into desired model
			var otpDetails OTP_details
			err := json.Unmarshal(e.Value, &otpDetails)

			fmt.Printf("%v", otpDetails)
			if err != nil {
				log.Fatal(err)
			}
			//-------------------------------------------------------------//

		case *kafka.Error:
			fmt.Printf("%s\n", e)
		}
	}
}
