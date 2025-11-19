package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	animals := []Animal{
		NewDog("Buddy", 5, time.Now()),
		NewMonkey("Cheeky", 3, time.Now()),
		NewLion("Leo", 8, time.Now()),
		NewTiger("Stripe", 7, time.Now()),
		NewCat("Whiskers", 4, time.Now()),
		NewHippo("Hank", 10, time.Now()),
		NewMonkeyGhost("Spooky", 50, time.Now()),
		NewLionGhost("Shadow", 100, time.Now()),
		NewTigerGhost("Phantom", 120, time.Now()),
		NewDog("Rocky", 2, time.Now()),
		NewCat("Smokey", 6, time.Now()),
		NewHippo("Bertha", 15, time.Now()),
		NewLion("Simba", 4, time.Now()),
		NewMonkey("George", 5, time.Now()),
		NewTiger("Rajah", 9, time.Now()),
	}

	for _, animal := range animals {
		animal.ProcessHealthCheck()
	}
}
