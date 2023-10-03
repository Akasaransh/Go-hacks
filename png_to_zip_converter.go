package main

import (
    "archive/zip"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    pngFileName := "example.png"
    zipFileName := "output.zip"
    zipFile, err := os.Create(zipFileName)
    if err != nil {
        fmt.Println("Error creating ZIP file:", err)
        return
    }
    defer zipFile.Close()
    zipWriter := zip.NewWriter(zipFile)
    pngData, err := ioutil.ReadFile(pngFileName)
    if err != nil {
        fmt.Println("Error reading PNG file:", err)
        return
    }
    pngFileWriter, err := zipWriter.Create(pngFileName)
    if err != nil {
        fmt.Println("Error creating ZIP file entry:", err)
        return
    }
    _, err = pngFileWriter.Write(pngData)
    if err != nil {
        fmt.Println("Error writing PNG data to ZIP file:", err)
        return
    }
    err = zipWriter.Close()
    if err != nil {
        fmt.Println("Error closing ZIP file:", err)
        return
    }

    fmt.Println("PNG file has been added to", zipFileName)
}
