package main

import (
    "fmt"
    "github.com/skip2/go-qrcode"
)

func main() {
    qrID := "Q12345"
    qrCode, err := qrcode.Encode(qrID, qrcode.Medium, 256)
    if err != nil {
        fmt.Println("Error generating QR code:", err)
        return
    }
    err = qrCode.Save(256, "qrCodeImage.png")
    if err != nil {
        fmt.Println("Error saving QR code as PNG:", err)
        return
    }

    fmt.Println("QR code saved as 'qrCodeImage.png'")
}
