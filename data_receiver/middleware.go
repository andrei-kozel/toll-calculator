package main

import (
	"time"

	"github.com/andrei-kozel/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddlware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"OBUID": data.OBUID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("producing to kafka")
	}(time.Now())

	return l.next.ProduceData(data)
}
