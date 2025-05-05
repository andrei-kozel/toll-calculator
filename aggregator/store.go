package main

import "github.com/andrei-kozel/toll-calculator/types"

type MemoryStrore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStrore {
	return &MemoryStrore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStrore) Insert(d types.Distance) error {
	m.data[d.OBUID] += d.Value
	return nil
}
