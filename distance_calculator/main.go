package main

import "log"

var kafkaTopic = "obudata"

func main() {
	calcService := NewCalcService()
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcService)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
