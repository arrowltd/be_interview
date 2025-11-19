package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- LionGhost ---

type LionGhost struct {
	name           string
	age            int
	firstDateAtZoo time.Time
}

func NewLionGhost(name string, age int, firstDateAtZoo time.Time) *LionGhost {
	return &LionGhost{name: name, age: age, firstDateAtZoo: firstDateAtZoo}
}

func (lg *LionGhost) Name() string {
	return lg.name
}

func (lg *LionGhost) Age() int {
	return lg.age
}

func (lg *LionGhost) FirstDateAtZoo() time.Time {
	return lg.firstDateAtZoo
}

func (lg *LionGhost) GiveOrder(order string) {
	fmt.Printf("%s received order: %s\n", lg.name, order)
}

func (lg *LionGhost) ProcessHealthCheck() {
	fmt.Printf("--- Starting health check for LionGhost: %s ---\n", lg.name)
	lg.GiveOrder("roar")
	lg.GiveOrder("claw")
	lg.GiveOrder("phase through wall")

	// Insert into fake database
	checkupNote := fmt.Sprintf("Random note %d", rand.Intn(100))
	nametag := fmt.Sprintf("LionGhost - %s - %d", lg.name, lg.age)
	fmt.Printf("DB INSERT: Name: %s, Age: %d, CheckupNote: %s, Nametag: %s\n", lg.name, lg.age, checkupNote, nametag)
	fmt.Printf("--- Finished health check for LionGhost: %s ---\n\n", lg.name)
}
