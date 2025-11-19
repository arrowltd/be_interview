package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- TigerGhost ---

type TigerGhost struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewTigerGhost(name string, age int, firstDateAtZoo time.Time) *TigerGhost {
	return &TigerGhost{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (tg *TigerGhost) Name() string {
	return tg.name
}

func (tg *TigerGhost) Age() int {
	return tg.age
}

func (tg *TigerGhost) FirstDateAtZoo() time.Time {
	return tg.firstDateAtZoo
}

func (tg *TigerGhost) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", tg.name, order)
}

func (tg *TigerGhost) ProcessHealthCheck() {
	fmt.Printf("--- Starting health check for TigerGhost: %s ---\n", tg.name)
	tg.GiveOrder("roar")
	tg.GiveOrder("claw")
	tg.GiveOrder("phase through wall")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("TigerGhost - %s - %d", tg.name, tg.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", tg.name, tg.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for TigerGhost: %s ---\n\n", tg.name)
}
