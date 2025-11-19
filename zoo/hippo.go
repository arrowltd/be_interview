package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Hippo ---

type Hippo struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewHippo(name string, age int, firstDateAtZoo time.Time) *Hippo {
	return &Hippo{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (h *Hippo) Name() string {
	return h.name
}

func (h *Hippo) Age() int {
	return h.age
}

func (h *Hippo) FirstDateAtZoo() time.Time {
	return h.firstDateAtZoo
}

func (h *Hippo) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", h.name, order)
}

func (h *Hippo) ProcessHealthCheck() {
	fmt.Printf("--- Starting health check for Hippo: %s ---\n", h.name)
	h.GiveOrder("swim")
	h.GiveOrder("eat")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("Hippo - %s - %d", h.name, h.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", h.name, h.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for Hippo: %s ---\n\n", h.name)
}
