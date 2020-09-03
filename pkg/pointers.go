package main

import (
	"encoding/json"
	"fmt"
	"go-pointers/pkg/entities"
	"math"
	"time"
)

// --- SIMULAMOS EL FLUJO DE UNA API NORMAL --- //
func serializationPointers() {
	// Unmarshal incoming requests
	var person entities.Pointers
	if err := json.Unmarshal(inputJSON, &person); err != nil {
		fmt.Println("error unmarshalling pointer struct", err)
	}

	// Do some changes to data
	// Note how we have to change how we use it
	*person.DNI = *person.DNI + 10
	*person.Money -= 0.5
	*person.FullName = *person.FullName + "MEAMEEEEEE"
	*person.Working = !*person.Working
	person.JoinDate.Add(time.Minute)

	// Marshal to send to other APIs
	_, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error unmarshalling pointer struct", err)
	}
}

func passingPointers() {
	person := entities.Values{
		DNI:      12345678,
		Money:    1200.53,
		FullName: "Person Name Jr.",
		Working:  true,
		JoinDate: time.Now(),
	}

	passPointer(&person, 0)
}

func passPointer(person *entities.Values, loop int) {
	if loop > 10 {
		return
	}

	// Do some changes to data
	person.DNI += 10
	person.Money -= 0.5
	person.FullName = person.FullName + "MEAMEEEEEE"
	person.Working = !person.Working
	person.JoinDate.Add(time.Minute)

	passPointer(person, loop+1)
}

func passingPointersLarge() {
	quijote := entities.QuijoteBook

	passPointersLarge(&quijote, 0)
}

func passPointersLarge(quijote *entities.Book, loop int) {
	if loop > 10 {
		return
	}

	passPointersLarge(quijote, loop+1)
}

type structPointerReciever struct{}

func (r *structPointerReciever) runPointerReceiver(s something) {
	for i := 0; i < (len(s.Nombre)+s.Edad)*100; i++ {
		cuadrado := i * i
		sqrt := math.Sqrt(float64(cuadrado))
		_ = sqrt * 2
	}
}
