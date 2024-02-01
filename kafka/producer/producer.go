package producer

import (
	"fmt"
	"os"

	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
	"github.com/segmentio/ksuid"
)

type OrderTransaction struct {
	ID     string
	Name   string
	Type   string
	Amount float64
}

func NewOrderTransaction() *OrderTransaction {
	return &OrderTransaction{}
}

var trans *OrderTransaction
func  SetOrderTransaction(o *OrderTransaction){
	trans = o
}

func getOrderTransaction() *OrderTransaction { 
	return trans 
}

func (o *OrderTransaction) SetApacheKafka(id, name, Ttype string, amount float64) {
	o.ID = id
	o.Name = name
	o.Type = Ttype
	o.Amount = amount
}

func (o *OrderTransaction) getApacheKafka() *OrderTransaction {
	return &OrderTransaction{
		ID:     o.ID,
		Name:   o.Name,
		Type:   o.Type,
		Amount: o.Amount,
	}
}

func ApacheKafkaProducer() error {
	config := sarama.NewConfig()

	config.Producer.Retry.Max = 5
	config.Producer.Transaction.Retry.Backoff = 10
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll

	config.Version = sarama.V2_7_2_0
	config.Net.MaxOpenRequests = 1

	txnID := ksuid.New().String()
	config.Producer.Transaction.ID = txnID
	config.Consumer.IsolationLevel = sarama.ReadCommitted

	// setting monitoring and logging
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Idempotent = true

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("FIRST_BROKER_URL"), os.Getenv("SECOND_BROKER_URL")}, config)
	if err != nil {
		log.Log.Panic(err)
	}

	defer func(producer sarama.SyncProducer) {
		if err := producer.Close(); err != nil {
			log.Log.Error(err)
		}
	}(producer)

	if err = producer.BeginTxn(); err != nil {
		log.Log.Error(err)
	}

	message := &sarama.ProducerMessage{
		Topic: assistance.TopicName,
		Value: sarama.StringEncoder(fmt.Sprintf("%+v", getOrderTransaction().getApacheKafka())),
	}

	if partition, offset, err := producer.SendMessage(message); err != nil {
		log.Log.Error(err)
		producer.AbortTxn()

	} else {
		fmt.Printf("Message sent to partition %d at offset %d", partition, offset)
		if err = producer.CommitTxn(); err != nil {
			log.Log.Error(err)
			producer.AbortTxn()

		}
	}

	return nil
}
