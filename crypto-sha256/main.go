package main

import (
	"crypto/sha256"

	"fmt"
)


func main() {

	crypto1 := sha256.Sum256([]byte("x"))
	crypto2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%T\n", crypto1, crypto2, crypto1)

}