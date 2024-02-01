package main

import (
	"os"

	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
)


func runApacheKafka() error {
	var config = sarama.NewConfig()
	config.Version = sarama.V2_7_2_0
	admin, err := sarama.NewClusterAdmin([]string{os.Getenv("FIRST_BROKER_URL"),os.Getenv("SECOND_BROKER_URL")}, config)
	if err != nil {
		log.Log.Panic("error creating kafka admid: " + err.Error())
	}

	defer func() {
		if err := admin.Close(); err != nil {
			log.Log.Error("error closing kafka admin: " + err.Error())
		}
	}()

	topicDetail := &sarama.TopicDetail{
		NumPartitions:     3,
		ReplicationFactor: 1,
	}
	

	topics, err := admin.ListTopics()
	if err != nil {
		log.Log.Panic("Error listing topics: ", err)
	}

	if _, ok := topics[assistance.TopicName]; ok {
		log.Log.Info("Like this topic exists already")
	} else {
		if err := admin.CreateTopic(assistance.TopicName, topicDetail, false); err != nil {
			log.Log.Panic("Failed to create topic apache Kafka: " + err.Error())
		}
	}

	return nil
}
