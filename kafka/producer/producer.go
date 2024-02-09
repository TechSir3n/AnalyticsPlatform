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

func (p *OrderAndProduct) SetDataProduct(id, name string, price float64, quantity int64) {
	p.Product = &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func (p *OrderAndProduct) GetDataProduct() (string, string, float64, int64) {
	return p.Product.ID, p.Product.Name, p.Product.Price, p.Product.Quantity
}

func (ot *OrderAndProduct) SetDataTrans(id, name, Ttype, time string, amount float64) {
	ot.Order = &OrderTransaction{
		ID:     id,
		Name:   name,
		Type:   Ttype,
		Time:   time,
		Amount: amount,
	}
}

func (ot *OrderAndProduct) GetDataTrans() (string, string, string, string, float64) {
	return ot.Order.ID, ot.Order.Name, ot.Order.Type, ot.Order.Time, ot.Order.Amount
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

	id, name, Ttype, time, amount := ot.GetDataTrans()
	id_product, name_product, price_product, quantity_product := ot.GetDataProduct()

	message := &sarama.ProducerMessage{
		Topic: assistance.TopicName,
		Value: sarama.StringEncoder(fmt.Sprintf("Orders: %s %s %s %s %f\n Products: %s %s %f %d", id, name,
			Ttype, time, amount, id_product, name_product, price_product, quantity_product)),
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
