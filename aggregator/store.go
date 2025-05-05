package main

import "github.com/andrei-kozel/toll-calculator/types"

type MemoryStrore struct{}

func NewMemoryStore() *MemoryStrore {
	return &MemoryStrore{}
}

func (m *MemoryStrore) Insert(d types.Distance) error {
	return nil
}
