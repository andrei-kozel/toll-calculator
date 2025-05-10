package main

import (
	"log"

	"github.com/andrei-kozel/toll-calculator/aggregator/client"
)

var kafkaTopic = "obudata"

func main() {
	calcService := NewCalcService()
	calcService = NewLogMiddleware(calcService)
	calculator := client.NewClient("http://127.0.0.1:3000/aggregate")
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcService, calculator)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
