package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/andrei-kozel/toll-calculator/types"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"OBUID": distance.OBUID,
			"value": distance.Value,
			"unix":  distance.Unix,
			"took":  time.Since(start),
		}).Info("aggregate distance")
	}(time.Now())

	return m.next.AggregateDistance(distance)
}
