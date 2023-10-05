package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	data := []byte("Sample data")

	seed := rand.Uint64()

	hash := universalHash(data, seed)

	fmt.Printf("Original Data: %s\n", data)
	fmt.Printf("Hash: %d\n", hash)
}

func universalHash(data []byte, seed uint64) uint64 {
	hash := fnv.New64a()

	hash.Write(data)

	hashSum := hash.Sum64()

	return hashSum ^ seed
}
