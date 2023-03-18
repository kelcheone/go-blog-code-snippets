package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

func main() {
	buffer := bytes.NewBuffer([]byte{120, 156, 10, 207, 47, 202, 206, 204, 75, 87, 40, 207, 44, 201, 80, 136, 202, 201, 76, 226,
		2, 4, 0, 0, 255, 255, 65, 134, 6, 121},
	)

	r, err := zlib.NewReader(buffer)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, r)
}
