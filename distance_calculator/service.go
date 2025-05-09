package main

import (
	"math"

	"github.com/andrei-kozel/toll-calculator/types"
)

type CalculatorService interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalcService struct {
	prevPoint []float64
}

func NewCalcService() CalculatorService {
	return &CalcService{}
}

func (s *CalcService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if len(s.prevPoint) > 0 {
		distance = calculateDistance(s.prevPoint[0], s.prevPoint[1], data.Lat, data.Long)
	}
	s.prevPoint = []float64{data.Lat, data.Long}
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}
