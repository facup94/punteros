package main

import (
	"testing"
)

// Serialization
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

// Pass by
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

func BenchmarkPassingNullables(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passingNullables()
	}
}

// Pass large struct
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

func BenchmarkPassingNullablesLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passingNullablesLarge()
	}
}

// Receivers
func BenchmarkValueReceiver(b *testing.B) {
	receiver := structReciever{}
	s := something{
		Nombre: "Facundo Parra",
		Edad:   25,
	}

	for n := 0; n < b.N; n++ {
		receiver.runValueReceiver(s)
	}
}
func BenchmarkPointerReceiver(b *testing.B) {
	receiver := &structReciever{}
	s := something{
		Nombre: "Facundo Parra",
		Edad:   25,
	}

	for n := 0; n < b.N; n++ {
		receiver.runPointerReceiver(s)
	}
}
