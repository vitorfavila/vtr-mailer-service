package tools

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func PrintStruct(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(b))
}

func GenerateRandId() int64 {
	// Defina o valor máximo para um int64.
	var max int64 = 1<<63 - 1

	// Gera um número aleatório int64 entre 0 e max.
	randomNumber := rand.Int63n(max)

	return randomNumber
}
