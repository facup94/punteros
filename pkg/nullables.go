package main

import (
	"encoding/json"
	"fmt"
	"go-pointers/pkg/entities"
	"time"
)

func serializationNullables() {
	// Unmarshal incoming requests
	var person entities.Nullables
	if err := json.Unmarshal(inputJSON, &person); err != nil {
		fmt.Println("error unmarshalling pointer struct", err)
	}

	//// Do some changes to data
	person.DNI.Value += 10
	person.Money.Value -= 0.5
	person.FullName.Value = person.FullName.Value + "MEAMEEEEEE"
	person.Working.Value = !person.Working.Value
	person.JoinDate.Value = person.JoinDate.Value.Add(time.Minute)

	//person.DNI = nullable.FromInt(person.DNI.Value + 10)
	//person.Money = nullable.FromFloat(person.Money.Value- 0.5)
	//person.FullName = nullable.FromString(person.FullName.Value + "MEAMEEEEEE")
	//person.Working = nullable.FromBool(!person.Working.Value)
	//person.JoinDate = nullable.FromTime(person.JoinDate.Value.Add(time.Minute))

	// Marshal to send to other APIs
	_, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error unmarshalling nullable struct", err)
	}
}
