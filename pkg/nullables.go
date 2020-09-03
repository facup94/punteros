package main

import (
	"encoding/json"
	"fmt"
	"go-pointers/pkg/entities"
	"time"

	"github.com/mercadolibre/fury_go-nullables/pkg/nullable"
)

func serializationNullables() {
	// Unmarshal incoming requests
	var person entities.Nullables
	if err := json.Unmarshal(inputJSON, &person); err != nil {
		fmt.Println("error unmarshalling pointer struct", err)
	}

	// Do some changes to data
	person.DNI.Value += 10
	person.Money.Value -= 0.5
	person.FullName.Value = person.FullName.Value + "MEAMEEEEEE"
	person.Working.Value = !person.Working.Value
	person.JoinDate.Value = person.JoinDate.Value.Add(time.Minute)

	// Marshal to send to other APIs
	_, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error unmarshalling nullable struct", err)
	}
}

func passingNullables() {
	person := entities.Nullables{
		DNI:      nullable.FromInt(12345678),
		Money:    nullable.FromFloat(1200.53),
		FullName: nullable.FromString("Person Name Jr."),
		Working:  nullable.FromBool(true),
		JoinDate: nullable.FromTime(time.Now()),
	}

	passNullable(person, 0)
}

func passNullable(person entities.Nullables, loop int) {
	if loop > 10 {
		return
	}

	// Do some changes to data
	person.DNI.Value += 10
	person.Money.Value -= 0.5
	person.FullName.Value = person.FullName.Value + "MEAMEEEEEE"
	person.Working.Value = !person.Working.Value
	person.JoinDate.Value.Add(time.Minute)

	passNullable(person, loop+1)
}

func passingNullablesLarge() {
	shakespeare := entities.ShakespeareBookNullable

	passNullableLarge(shakespeare, 0)
}

func passNullableLarge(shakespeare entities.BookNullable, loop int) {
	if loop > 10 {
		return
	}

	passNullableLarge(shakespeare, loop+1)
}
