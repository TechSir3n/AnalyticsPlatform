package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
	"os"
)

func apacheKafkaConsumer() error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.IsolationLevel = sarama.ReadCommitted

	consumer, err := sarama.NewConsumer([]string{os.Getenv("FIRST_BROKER_URL"),os.Getenv("SECOND_BROKER_URL")}, config)
	if err != nil {
		log.Log.Panic(err)
	}

	defer func(consumer sarama.Consumer) {
		if err := consumer.Close(); err != nil {
			log.Log.Error(err)
		}
	}(consumer)

	consumerPartition, err := consumer.ConsumePartition(assistance.TopicName, 0, sarama.OffsetOldest)
	if err != nil {
		log.Log.Panic(err)
	}

	defer func(consumerPartotion sarama.PartitionConsumer) {
		if err := consumerPartition.Close(); err != nil {
			log.Log.Error(err)
		}
	}(consumerPartition)

	for {
		select {
		case msg := <-consumerPartition.Messages():
			fmt.Println(string(msg.Value))
		case err := <-consumerPartition.Errors():
			log.Log.Error(err)
		}
	}
}
