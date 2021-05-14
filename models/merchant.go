package models

import "sync"

type Merchant struct {
	sync.RWMutex
	Name            string
	Email           string
	DiscountPercent float64
}
