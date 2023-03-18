package main

import (
	"bytes"
	"compress/zlib"
	"image"
	"image/png"
	"os"
)

// Task: Compress an image using zlib and save the compressed image to a file.

func main() {
	// Open the image file.
	file, err := os.Open("image1.png")
	if err != nil {
		panic(err)
	}

	// Decode the image.
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Create a new buffer.
	var buf bytes.Buffer

	// Create a new zlib writer.
	w := zlib.NewWriter(&buf)

	// Encode the image to the buffer.
	err = png.Encode(w, img)
	if err != nil {
		panic(err)
	}

	// save the compressed image to a file
	err = os.WriteFile("compressed.png", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
