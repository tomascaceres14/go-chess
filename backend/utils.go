package main

import (
	"fmt"
	"math/rand"
)

func randomBool() bool {
	return rand.Intn(2) == 0
}

func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}
