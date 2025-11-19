package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Ethereal ---

type Ethereal struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewEthereal(name string, age int, firstDateAtZoo time.Time) *Ethereal {
	return &Ethereal{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (e *Ethereal) Name() string {
	return e.name
}

func (e *Ethereal) Age() int {
	return e.age
}

func (e *Ethereal) FirstDateAtZoo() time.Time {
	return e.firstDateAtZoo
}

func (e *Ethereal) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", e.name, order)
}

func (e *Ethereal) ProcessHealthCheck(history *MedicalHistoryObject) {
	fmt.Printf("--- Starting health check for Ethereal: %s ---\n", e.name)
	e.GiveOrder("pass through wall")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("Ethereal - %s - %d", e.name, e.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", e.name, e.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for Ethereal: %s ---\n\n", e.name)
}
