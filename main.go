package main

import (
	"awesomeProject/algorithms"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Grettings: ")
	log.SetFlags(0)

	messsage, err := algorithms.Message()

	if err != nil {
		log.Fatalf("Mensagem invalida")
	}
	fmt.Println(messsage)
}
