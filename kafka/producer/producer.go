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
	Time   string
	Amount float64
}

type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int64
}

type OrderAndProduct struct {
	Order   *OrderTransaction
	Product *Product
}

func (p *Product) SetData(id, name string, price float64, quantity int64) {
	p.ID = id
	p.Name = name
	p.Price = price
	p.Quantity = quantity
}

func (p *Product) GetData() (string, string, float64, int64) {
	return p.ID, p.Name, p.Price, p.Quantity
}

func (ot *OrderTransaction) SetData(id, name, Ttype, time string, amount float64) {
	ot.ID = id
	ot.Name = name
	ot.Type = Ttype
	ot.Time = time
	ot.Amount = amount
}

func (ot *OrderTransaction) GetData() (string, string, string, string, float64) {
	return ot.ID, ot.Name, ot.Type, ot.Time, ot.Amount
}

func SetObject(_order *OrderAndProduct) {
	_order.ApacheKafkaProducerRun()
}

func (ot *OrderAndProduct) ApacheKafkaProducerRun() error {
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

	id, name, Ttype, time, amount := ot.Order.GetData()
	id_product, name_product, price_product, quantity_product := ot.Product.GetData()

	message := &sarama.ProducerMessage{
		Topic: assistance.TopicName,
		Value: sarama.StringEncoder(fmt.Sprintf("Orders: %s %s %s %s %f\n Products: %s %s %f %d ",
			id, name, Ttype, time, amount, id_product, name_product, price_product, quantity_product)),
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
