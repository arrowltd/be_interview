package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Cat ---

type Cat struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewCat(name string, age int, firstDateAtZoo time.Time) *Cat {
	return &Cat{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (c *Cat) Name() string {
	return c.name
}

func (c *Cat) Age() int {
	return c.age
}

func (c *Cat) FirstDateAtZoo() time.Time {
	return c.firstDateAtZoo
}

func (c *Cat) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", c.name, order)
}

func (c *Cat) ProcessHealthCheck(history *MedicalHistoryObject) {
	fmt.Printf("--- Starting health check for Cat: %s ---\n", c.name)
	c.GiveOrder("mew")
	c.GiveOrder("claw")
	c.GiveOrder("jump")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("Cat - %s - %d", c.name, c.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", c.name, c.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for Cat: %s ---\n\n", c.name)
}
