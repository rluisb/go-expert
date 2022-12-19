package main

import (
	"curso-go/matematica"
	"github.com/google/uuid"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado: ", s)
	fmt.Println(uuid.New())
}
