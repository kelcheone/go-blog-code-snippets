package main

// Task: Decompress an image using zlib and save the decompressed image to a file.

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	// open the compressed image file
	file, err := os.Open("compressed.png")
	if err != nil {
		log.Fatal(err)
	}

	// read the compressed image file
	buf := new(bytes.Buffer)

	// create a new zlib reader
	r, err := zlib.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// copy the compressed image to the buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		log.Fatal(err)
	}

	// decode the image
	img, err := png.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}

	// save the decompressed image to a file
	err = os.WriteFile("decompressed.png", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// print the image bounds
	fmt.Println(img.Bounds())
}
