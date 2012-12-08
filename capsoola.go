package main

import (
	"reflect"
	"appengine"
	"appengine/datastore"
)

// define Map

type Map map[string]interface{}

func (m Map) Load(c <-chan datastore.Property) error {
	for p := range c {
		if p.Multiple {
			value := reflect.ValueOf(m[p.Name])
			if value.Kind() != reflect.Slice {
				m[p.Name] = []interface{}{p.Value}
			} else {
				m[p.Name] = append(m[p.Name].([]interface{}), p.Value)
			}
		} else {
			m[p.Name] = p.Value	
		}
	}
	return nil
}

func (m Map) Save(c chan<- datastore.Property) error {
	defer close(c)
	for k, v := range m {
		c <- datastore.Property {
			Name: k,
			Value: v,
		}
	}
	return nil
}