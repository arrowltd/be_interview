package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- MonkeyGhost ---

type MonkeyGhost struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewMonkeyGhost(name string, age int, firstDateAtZoo time.Time) *MonkeyGhost {
	return &MonkeyGhost{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (mg *MonkeyGhost) Name() string {
	return mg.name
}

func (mg *MonkeyGhost) Age() int {
	return mg.age
}

func (mg *MonkeyGhost) FirstDateAtZoo() time.Time {
	return mg.firstDateAtZoo
}

func (mg *MonkeyGhost) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", mg.name, order)
}

func (mg *MonkeyGhost) ProcessHealthCheck(history *MedicalHistoryObject) {
	fmt.Printf("--- Starting health check for MonkeyGhost: %s ---\n", mg.name)
	mg.GiveOrder("climb")
	mg.GiveOrder("jump")
	mg.GiveOrder("phase through wall")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("MonkeyGhost - %s - %d", mg.name, mg.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", mg.name, mg.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for MonkeyGhost: %s ---\n\n", mg.name)
}
