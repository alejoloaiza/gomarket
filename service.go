package main

import (
	"encoding/json"

	"github.com/pkg/errors"
)

/*
func (o *Order) Parse(line interface{}) (bool, error) {
	s := reflect.ValueOf(line)
	if s.Kind() == reflect.Slice {

	}
}
*/

func (o *Order) toBytes() []byte {
	bytes, _ := json.Marshal(o)
	return bytes
}
func (o *Order) toString() string {
	str, _ := json.Marshal(o)
	return string(str)
}
func (o *Order) validate() error {
	if o.Amount == int64(0) || o.Asset == "" || o.ID == "" || o.Price == int64(0) || o.Type == "" {
		return errors.New("Cannot insert order, Mandatory field validation")
	}
	return nil
}
