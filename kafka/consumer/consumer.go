package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	"github.com/TechSir3n/analytics-platform/clickhouse"
	"github.com/TechSir3n/analytics-platform/kafka/consumer/models"
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

	db_transaction, err := clickhouse.NewClickHouseClient("Transaction")
	if err != nil {
		log.Log.Error("[Product Database]: ", err)
	}

	db_product, err := clickhouse.NewClickHouseClient("Product")
	if err != nil {
		log.Log.Error("[Transaction Database]: ", err)
	}

	for {
		select {
		case msg := <-partitionTransaction.Messages():
			var transaction models.Transaction
			if err := json.Unmarshal(msg.Value, &transaction); err != nil {
				log.Log.Error("Failed unmarshal struct [transaction]: ", err)
			}

			if err = db_transaction.InsertData(transaction.Order_id, 0, transaction.Customer_name,
				transaction.Transaction_type, transaction.Amount, 0.0, 0.0, 0.0, transaction.Transaction_date); err != nil {
				log.Log.Error("Failed insert data to transactionDB: ", err)
			}

		case msg := <-partitionProduct.Messages():
			var product models.Product
			if err := json.Unmarshal(msg.Value, &product); err != nil {
				log.Log.Error("Failed unmarshal struct [product]: ", err)
			}

			if err = db_product.InsertData(product.Product_id, product.Quantity, product.Product_name, "", 0.0, product.Price, 0.0, 0.0, ""); err != nil {
				log.Log.Error("Failed insert data to productDB: ", err)
			}

		case err := <-partitionTransaction.Errors():
			log.Log.Error(err)
		case err := <-partitionProduct.Errors():
			log.Log.Error(err)
		}
	}
}
