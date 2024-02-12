package producer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
)

func SendApacheBrokerTrans(id, name, Ttype, time string, amount float64, producer sarama.SyncProducer) {
	if err := producer.BeginTxn(); err != nil {
		log.Log.Error(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: assistance.TopicTransaction,
		Value: sarama.StringEncoder(fmt.Sprintf("Transaction: %s %s %s %s %f", id, name, Ttype, time, amount)),
	}

	if _, _, err := producer.SendMessage(msg); err != nil {
		log.Log.Errorf("Failed to send message: %s", err)
		producer.AbortTxn()
	} else {
		if err := producer.CommitTxn(); err != nil {
			log.Log.Errorf("Commit transaction error: %v", err)
			producer.AbortTxn()
		}
	}
}

func SendApacheBrokerProduct(id, name string, price float64, quantity int64, producer sarama.SyncProducer) {
	if err := producer.BeginTxn(); err != nil {
		log.Log.Error(err)
	}
	msg := &sarama.ProducerMessage{
		Topic: assistance.TopicProduct,
		Value: sarama.StringEncoder(fmt.Sprintf("Product: %s %s %f %d", id, name, price, quantity)),
	}

	if _, _, err := producer.SendMessage(msg); err != nil {
		log.Log.Errorf("Failed to send message: %s", err)
	} else {
		if err := producer.CommitTxn(); err != nil {
			log.Log.Errorf("Commit transaction error: %v", err)
			producer.AbortTxn()
		}
	}
}
