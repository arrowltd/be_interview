package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Monkey ---

type Monkey struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewMonkey(name string, age int, firstDateAtZoo time.Time) *Monkey {
	return &Monkey{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (m *Monkey) Name() string {
	return m.name
}

func (m *Monkey) Age() int {
	return m.age
}

func (m *Monkey) FirstDateAtZoo() time.Time {
	return m.firstDateAtZoo
}

func (m *Monkey) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", m.name, order)
}

func (m *Monkey) ProcessHealthCheck(history *MedicalHistoryObject) {
	fmt.Printf("--- Starting health check for Monkey: %s ---\n", m.name)
	m.GiveOrder("climb")
	m.GiveOrder("jump")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("Monkey - %s - %d", m.name, m.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", m.name, m.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for Monkey: %s ---\n\n", m.name)
}
