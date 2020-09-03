package main

import (
	"testing"
)

func BenchmarkSerializationValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serializationValues()
	}
}

func BenchmarkSerializationPointers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serializationPointers()
	}
}

func BenchmarkSerializationNullables(b *testing.B) {
	for n := 0; n < b.N; n++ {
		serializationNullables()
	}
}

func BenchmarkPassingValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passingValues()
	}
}

func BenchmarkPassingPointers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passingPointers()
	}
}

func BenchmarkPassingValuesLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passingValuesLarge()
	}
}

func BenchmarkPassingPointersLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passingPointersLarge()
	}
}

func BenchmarkValueReceiver(b *testing.B) {
	receiver := structValueReciever{}
	s := something{
		Nombre: "Facundo Parra",
		Edad:   25,
	}

	for n := 0; n < b.N; n++ {
		receiver.runValueReceiver(s)
	}
}
func BenchmarkPointerReceiver(b *testing.B) {
	receiver := &structPointerReciever{}
	s := something{
		Nombre: "Facundo Parra",
		Edad:   25,
	}

	for n := 0; n < b.N; n++ {
		receiver.runPointerReceiver(s)
	}
}
