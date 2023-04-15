package kafka_producer

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OTP_Payload struct {
	Otp   string
	Email string
}

type OTP_Producer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func New_OTP_Producer(p *kafka.Producer, topic string) *OTP_Producer {
	return &OTP_Producer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 100),
	}
}

func (op *OTP_Producer) Send_OTP(otpPayload OTP_Payload) error {
	payload, err := json.Marshal(otpPayload)

	if err != nil {
		log.Fatal(err)
	}

	err = op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny},
		Value: payload},
		op.deliverych,
	)

	if err != nil {
		log.Fatal(err)
	}

	<-op.deliverych
	return nil

}

func ProduceOTP(otpPayload OTP_Payload) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "unique1222",
		"acks":              "all",
	})

	if err != nil {
		return err
	}

	op := New_OTP_Producer(p, "OTP")
	if err := op.Send_OTP(otpPayload); err != nil {
		return err
	}

	log.Println("product sent......")

	return nil
}
