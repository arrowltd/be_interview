package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type MedicalHistoryObject struct {
	Note   string
	Status string
}

func fetchMedicalHistory(animal Animal) *MedicalHistoryObject {
	fmt.Println("--- Calling 3rd party service to fetch medical history ---")

	// Manually create nametag
	className := reflect.TypeOf(animal).Elem().Name()
	nametag := fmt.Sprintf("%s - %s - %d", className, animal.Name(), animal.Age())

	fmt.Printf("Sending request for nametag: %s\n", nametag)

	// Fake receiving response
	history := &MedicalHistoryObject{
		Note:   fmt.Sprintf("Random medical note %d", rand.Intn(100)),
		Status: "alive",
	}

	fmt.Printf("Received response: %+v\n", history)
	fmt.Println("--- End of 3rd party call ---")

	return history
}
