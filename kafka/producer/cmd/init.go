package init

import (
	"github.com/IBM/sarama"
	_ "github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
	"github.com/segmentio/ksuid"
	"os"
)

var Producer sarama.SyncProducer

func init() {
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

	var err error
	Producer, err = sarama.NewSyncProducer([]string{os.Getenv("FIRST_BROKER_URL"), os.Getenv("SECOND_BROKER_URL")}, config)
	if err != nil {
		log.Log.Panic(err)
	}
}

func CloseProducer() {
	if err := Producer.Close(); err != nil {
		log.Log.Error(err)
	}
}
