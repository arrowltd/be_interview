package main

import "time"

// Animal represents an animal in the zoo.
type Animal interface {
	Name() string
	Age() int
	FirstDateAtZoo() time.Time
	GiveOrder(order string)
	ProcessHealthCheck()
}
