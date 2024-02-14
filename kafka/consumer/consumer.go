package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	// processor "github.com/TechSir3n/analytics-platform/data_processing"
	log "github.com/TechSir3n/analytics-platform/logging"
	"os"
)

func apacheKafkaConsumer() error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.IsolationLevel = sarama.ReadCommitted

	consumer, err := sarama.NewConsumer([]string{os.Getenv("FIRST_BROKER_URL"), os.Getenv("SECOND_BROKER_URL")}, config)
	if err != nil {
		log.Log.Panic(err)
	}

	defer func(consumer sarama.Consumer) {
		if err := consumer.Close(); err != nil {
			log.Log.Error(err)
		}
	}(consumer)

	partitionTransaction, err := consumer.ConsumePartition(assistance.TopicTransaction, 0, sarama.OffsetOldest)
	if err != nil {
		log.Log.Panic(err)
	}

	partitionProduct, err := consumer.ConsumePartition(assistance.TopicProduct, 0, sarama.OffsetNewest)
	if err != nil {
		log.Log.Panic(err)
	}

	defer func(partitionTrans sarama.PartitionConsumer, partitionProduct sarama.PartitionConsumer) {
		if err := partitionTrans.Close(); err != nil {
			log.Log.Error(err)
		}
		if err := partitionProduct.Close(); err != nil {
			log.Log.Error(err)
		}
	}(partitionTransaction, partitionProduct)

	for {
		select {
		case msg := <-partitionTransaction.Messages():
			fmt.Println("Topic[Transaction]: ", string(msg.Value))
		case msg := <-partitionProduct.Messages():
			fmt.Println("Topic[Product]: ", string(msg.Value))
		case err := <-partitionTransaction.Errors():
			log.Log.Error(err)
		case err := <-partitionProduct.Errors():
			log.Log.Error(err)
		}
	}
}
