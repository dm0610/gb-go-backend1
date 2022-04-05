package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("random number: %d\n", n.Int64())
}
