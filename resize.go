package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	// Define flags
	width := flag.Uint("width", 500, "Width of the resized image")
	height := flag.Uint("height", 500, "Height of the resized image")
	convertFormat := flag.String("convert", "", "Target format for conversion (jpg or png)")

	// Custom argument parsing
	var inputFilePath string
	args := os.Args[1:] // Skip program name

	// Process arguments manually
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--width", "-width":
			if i+1 < len(args) {
				_, err := fmt.Sscanf(args[i+1], "%d", width)
				if err != nil {
					log.Fatal("Invalid width value")
				}
				i++
			}
		case "--height", "-height":
			if i+1 < len(args) {
				_, err := fmt.Sscanf(args[i+1], "%d", height)
				if err != nil {
					log.Fatal("Invalid height value")
				}
				i++
			}
		case "--convert", "-convert":
			if i+1 < len(args) {
				*convertFormat = args[i+1]
				if *convertFormat != "jpg" && *convertFormat != "png" {
					log.Fatal("Error: -convert flag must be 'jpg' or 'png'")
				}
				i++
			}
		default:
			// If it's not a flag, treat it as the input filename
			if inputFilePath == "" {
				inputFilePath = args[i]
			} else {
				log.Fatal("Error: Multiple input files not supported")
			}
		}
	}

	// Ensure we have an input file
	if inputFilePath == "" {
		log.Fatal("Error: Missing image file argument")
	}

	// Extract filename info
	inputName := inputFilePath[:len(inputFilePath)-len(filepath.Ext(inputFilePath))]
	inputExt := filepath.Ext(inputFilePath)
	outName := inputName + fmt.Sprintf("_%dx%d", *width, *height)

	// Open the input file
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Decode the image
	img, format, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Resize the image
	resizedImg := resize.Resize(*width, *height, img, resize.Lanczos3)

	// Determine output filename & extension
	var finalExt string
	if *convertFormat == "jpg" {
		finalExt = ".jpg"
	} else if *convertFormat == "png" {
		finalExt = ".png"
	} else {
		finalExt = inputExt // Keep original
	}

	outName = inputName + fmt.Sprintf("_%dx%d%s", *width, *height, finalExt)

	// Debugging Information
	fmt.Printf("Input File: %s\n", inputFilePath)
	fmt.Printf("Detected Format: %s\n", format)
	fmt.Printf("Output File: %s\n", outName)
	fmt.Printf("Requested Conversion: %s\n", *convertFormat)

	// Create the output file
	outputFile, err := os.Create(outName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Encode image in the correct format
	if *convertFormat == "jpg" {
		err = jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 90})
	} else if *convertFormat == "png" {
		err = png.Encode(outputFile, resizedImg)
	} else {
		// Default: use detected format
		if format == "png" {
			err = png.Encode(outputFile, resizedImg)
		} else if format == "jpeg" {
			err = jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 90})
		} else {
			log.Fatal("Unsupported image format")
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Converted and resized image saved as %s\n", outName)
}
