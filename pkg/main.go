package main

import "fmt"

func main() {
	fmt.Println("empty main")
	serializationNullables()
}

type something struct {
	Nombre string
	Edad   int
}
