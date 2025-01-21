package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Directory containing PNG files
	sourceDir := "/mnt/c/Users/ASUS/Downloads/FOTO FOR APPS"
	targetDir := "/mnt/c/Users/ASUS/Downloads/prod"

	// Ensure the output directory exists
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create output directory: %v\n", err)
		return
	}

	// Walk through the source directory to find PNG files
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current file is a PNG image
		if !info.IsDir() && (filepath.Ext(path) == ".png" || filepath.Ext(path) == ".jpg") {
			// Generate the output file path while maintaining the directory structure
			relPath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				fmt.Printf("Failed to get relative path for %s: %v\n", path, err)
				return err
			}
			outputFilePath := filepath.Join(targetDir, relPath[:len(relPath)-len(filepath.Ext(relPath))]+".webp")

			// Create subdirectories in the output directory if needed
			if err := os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
				fmt.Printf("Failed to create subdirectory for %s: %v\n", outputFilePath, err)
				return err
			}

			// Command to convert PNG to WebP
			cmd := exec.Command("cwebp", path, "-o", outputFilePath)

			// Run the cwebp command
			if err := cmd.Run(); err != nil {
				fmt.Printf("Failed to convert %s: %v\n", path, err)
				return err
			}

			fmt.Printf("Converted: %s -> %s\n", path, outputFilePath)
		}
		
		return nil
	})

	if err != nil {
		fmt.Printf("Error processing files: %v\n", err)
		return
	}

	fmt.Println("All PNG images successfully converted to WebP.")
}
