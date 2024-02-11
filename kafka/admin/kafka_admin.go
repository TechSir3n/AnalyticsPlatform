package main

import (
	"os"

	"github.com/IBM/sarama"
	"github.com/TechSir3n/analytics-platform/assistance"
	log "github.com/TechSir3n/analytics-platform/logging"
)

func createAdminApacheKafka() error {
	var config = sarama.NewConfig()
	config.Version = sarama.V2_7_2_0
	admin, err := sarama.NewClusterAdmin([]string{os.Getenv("FIRST_BROKER_URL"), os.Getenv("SECOND_BROKER_URL")}, config)
	if err != nil {
		return err
	}

	defer func() error {
		if err := admin.Close(); err != nil {
			return err
		}
		return nil
	}()

	topicDetail := &sarama.TopicDetail{
		NumPartitions:     3,
		ReplicationFactor: 1,
		//ConfigEntries:     map[string]*string{},
	}

	topics, err := admin.ListTopics()
	if err != nil {
		return err
	}

	if _, ok := topics[assistance.TopicTransaction]; ok {
		log.Log.Info("this topic exists already[TopicTransaction]")
	}

	if _, ok := topics[assistance.TopicProduct]; ok {
		log.Log.Info("this topic exists already[TopicProduct]")
	}

	if err := admin.CreateTopic(assistance.TopicTransaction, topicDetail, false); err != nil {
		return err
	}

	if err := admin.CreateTopic(assistance.TopicProduct, topicDetail, false); err != nil {
		return err
	}

	return nil
}

