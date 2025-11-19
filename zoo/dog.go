package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Dog ---

type Dog struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewDog(name string, age int, firstDateAtZoo time.Time) *Dog {
	return &Dog{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (d *Dog) Name() string {
	return d.name
}

func (d *Dog) Age() int {
	return d.age
}

func (d *Dog) FirstDateAtZoo() time.Time {
	return d.firstDateAtZoo
}

func (d *Dog) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", d.name, order)
}

func (d *Dog) ProcessHealthCheck() {
	fmt.Printf("--- Starting health check for Dog: %s ---\n", d.name)
	d.GiveOrder("lift paw")
	d.GiveOrder("bark")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("Dog - %s - %d", d.name, d.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", d.name, d.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for Dog: %s ---\n\n", d.name)
}
