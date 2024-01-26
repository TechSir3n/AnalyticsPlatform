package main

import (
	"fmt"
	"os"
	"time"

	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
	"github.com/segmentio/ksuid"
)

func apacheKafkaProducer() error {
	config := sarama.NewConfig()

	config.Producer.Retry.Max = 5
	config.Producer.Transaction.Retry.Backoff = 10
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	config.Version = sarama.V2_7_2_0
	config.Net.MaxOpenRequests = 1

	txnID := ksuid.New().String()
	config.Producer.Transaction.ID = txnID
	config.Consumer.IsolationLevel = sarama.ReadUncommitted // by default

	// setting monitoring and logging
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Idempotent = true

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("FIRST_BROKER_URL")}, config)
	if err != nil {
		log.Log.Panic(err)
	}

	defer func(producer sarama.SyncProducer) {
		if err := producer.Close(); err != nil {
			log.Log.Error(err)
		}
	}(producer)


	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	var transaction assistance.Transaction

	for range ticker.C {
		transaction = assistance.GenerateTransaction()

		if err = producer.BeginTxn(); err != nil {
			log.Log.Error(err)
		}
		
		message := &sarama.ProducerMessage{
			Topic: assistance.TopicName,
			Value: sarama.StringEncoder(fmt.Sprintf("%+v", transaction)),
		}

		if partition, offset, err := producer.SendMessage(message); err != nil {
			log.Log.Error(err)
			producer.AbortTxn()
			continue
		} else {
			fmt.Printf("Message sent to partition %d at offset %d", partition, offset)
			if err = producer.CommitTxn(); err != nil {
				log.Log.Error(err)
				producer.AbortTxn()
				continue
			}
		}
		time.Sleep(time.Second)
	}
	return nil
}
