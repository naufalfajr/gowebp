package main

import (
    "fmt"
    "os/exec"
)

func single() {
	home := "/mnt/c/Users/ASUS/Downloads/FOTO FOR APPS/"
    inputFile := home+"DSC08325.png"
    outputFile := home+"output.webp"
    cmd := exec.Command("cwebp", inputFile, "-o", outputFile)

    // Run the command
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Failed to run cwebp: %v\n", err)
        return
    }

    fmt.Println("Image successfully converted to WebP.")
}