package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Tiger ---

type Tiger struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewTiger(name string, age int, firstDateAtZoo time.Time) *Tiger {
	return &Tiger{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (t *Tiger) Name() string {
	return t.name
}

func (t *Tiger) Age() int {
	return t.age
}

func (t *Tiger) FirstDateAtZoo() time.Time {
	return t.firstDateAtZoo
}

func (t *Tiger) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", t.name, order)
}

func (t *Tiger) ProcessHealthCheck(history *MedicalHistoryObject) {
	fmt.Printf("--- Starting health check for Tiger: %s ---\n", t.name)
	t.GiveOrder("roar")
	t.GiveOrder("claw")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("Tiger - %s - %d", t.name, t.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", t.name, t.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for Tiger: %s ---\n\n", t.name)
}
