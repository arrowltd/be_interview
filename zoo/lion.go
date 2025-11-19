package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Lion ---

type Lion struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewLion(name string, age int, firstDateAtZoo time.Time) *Lion {
	return &Lion{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (l *Lion) Name() string {
	return l.name
}

func (l *Lion) Age() int {
	return l.age
}

func (l *Lion) FirstDateAtZoo() time.Time {
	return l.firstDateAtZoo
}

func (l *Lion) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", l.name, order)
}

func (l *Lion) ProcessHealthCheck(history *MedicalHistoryObject) {
	fmt.Printf("--- Starting health check for Lion: %s ---\n", l.name)
	l.GiveOrder("roar")
	l.GiveOrder("claw")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("Lion - %s - %d", l.name, l.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", l.name, l.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for Lion: %s ---\n\n", l.name)
}
