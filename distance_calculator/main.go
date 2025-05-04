package main

import "log"

var kafkaTopic = "obudata"

func main() {
	calcService := NewCalcService()
	calcService = NewLogMiddleware(calcService)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcService)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
