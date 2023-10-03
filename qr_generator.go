package main

import (
    "fmt"
    "math/rand"
    "time"
)

func generateUniqueQRID() string {
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(1000000)
    qrID := fmt.Sprintf("Q%06d", randomNumber)
    return qrID
}

func main() {
    qrID := generateUniqueQRID()
    fmt.Println("Generated QR ID:", qrID)
}
