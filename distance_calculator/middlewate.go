package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/andrei-kozel/toll-calculator/types"
)

type LogMiddleware struct {
	next CalculatorService
}

func NewLogMiddleware(next CalculatorService) CalculatorService {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"err":  err,
			"took": time.Since(start),
			"dist": dist,
		}).Info("calculating result")
	}(time.Now())
	return l.next.CalculateDistance(data)
}
