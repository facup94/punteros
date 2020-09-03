package main

import (
	"encoding/json"
	"fmt"
	"go-pointers/pkg/entities"
	"math"
	"time"
)

var inputJSON = []byte("{\"dni\": 12345678, \"money\": 1200.53, \"full_name\": \"Person Name Jr.\", \"working\": true, \"join_date\": \"2018-09-22T12:42:31Z\" }")

// --- SIMULAMOS EL FLUJO DE UNA API NORMAL --- //
func serializationValues() {
	// Unmarshal incoming requests
	var person entities.Values
	if err := json.Unmarshal(inputJSON, &person); err != nil {
		fmt.Println("error unmarshalling value struct", err)
	}

	// Do some changes to data
	person.DNI += 10
	person.Money -= 0.5
	person.FullName = person.FullName + "MEAMEEEEEE"
	person.Working = !person.Working
	person.JoinDate.Add(time.Minute)

	// Marshal to send to other APIs
	_, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error unmarshalling value struct", err)
	}
}

func passingValues() {
	person := entities.Values{
		DNI:      12345678,
		Money:    1200.53,
		FullName: "Person Name Jr.",
		Working:  true,
		JoinDate: time.Now(),
	}

	passValue(person, 0)
}

func passValue(person entities.Values, loop int) {
	if loop > 10 {
		return
	}

	// Do some changes to data
	person.DNI += 10
	person.Money -= 0.5
	person.FullName = person.FullName + "MEAMEEEEEE"
	person.Working = !person.Working
	person.JoinDate.Add(time.Minute)

	passValue(person, loop+1)
}

func passingValuesLarge() {
	quijote := entities.ShakespeareBook

	passValueLarge(quijote, 0)
}

func passValueLarge(quijote entities.Book, loop int) {
	if loop > 10 {
		return
	}

	passValueLarge(quijote, loop+1)
}

type structValueReciever struct{}

func (r structValueReciever) runValueReceiver(s something) {
	for i := 0; i < (len(s.Nombre)+s.Edad)*100; i++ {
		cuadrado := i * i
		sqrt := math.Sqrt(float64(cuadrado))
		_ = sqrt * 2
	}
}
