package main

import (
	"time"
)

// Order is the structs that defines an article for the API.
type Order struct {
	ID         string    `json:"id"`
	Asset      string    `json:"asset"`
	Amount     int64     `json:"amount"`
	EntryDate  time.Time `json:"entrydate"`
	ExpiryDate time.Time `json:"expirydate"`
	Type       string    `json:"type"`
	Price      int64     `json:"price"`
}

// Configuration is used to store all the configuration fields from the JSON config
type Configuration struct {
	ValidCommands   []string
	MinWordsCommand int
}
