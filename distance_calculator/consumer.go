package main

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"

	"github.com/andrei-kozel/toll-calculator/aggregator/client"
	"github.com/andrei-kozel/toll-calculator/types"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorService
	aggClient   *client.Client
}

func NewKafkaConsumer(topic string, svc CalculatorService, aggClient *client.Client) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
		aggClient:   aggClient,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("kafka consumer started")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("kafka consume error: %s", err)
			continue
		}
		var data types.OBUData
		if err = json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON deserialization error: %s", err)
			continue
		}

		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("cacl error: %s", err)
			continue
		}
		req := types.Distance{
			Value: distance,
			Unix:  time.Now().UnixNano(),
			OBUID: data.OBUID,
		}
		if err := c.aggClient.AggregateInvoice(req); err != nil {
			logrus.Errorf("agregate error: %v", err)
			continue
		}
	}
}
