package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
	"os"
)

func apacheKafkaProducer() error {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.Transaction.Retry.Backoff = 10
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Version = sarama.V2_7_2_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Net.MaxOpenRequests = 1

	// setting monitoring and logging
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Idempotent = true
															
	producer, err := sarama.NewSyncProducer([]string{os.Getenv("FIRST_BROKER_URL")}, config)
	if err != nil {
		log.Log.Panic("Failed to initialyize Apache-Producer: ", err.Error())
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Log.Error("Failed to close Producer: ", err)
		}
	}()

	transaction := assistance.GenerateTransaction()
	message := &sarama.ProducerMessage{
		Topic: assistance.TopicName,
		Value: sarama.StringEncoder(fmt.Sprintf("%+v", transaction)),
	}

	if _, _, err = producer.SendMessage(message); err != nil {
		log.Log.Panic(err)
	}

	fmt.Println("Success sent message")

	return nil
}
