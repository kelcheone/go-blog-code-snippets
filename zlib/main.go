package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a buffer to write our uncompressed data to.
	var b bytes.Buffer

	// Create a new zlib writer.
	w := zlib.NewWriter(&b)

	// Write some uncompressed data to the zlib writer.
	// Any source that implements the io.Writer interface can be used.
	w.Write([]byte("Working with Zlib\n"))

	// Write some more uncompressed data to the zlib writer.
	// fmt.Fprintln(w, "another line here")

	// Close the zlib writer.
	// This is important to flush the bytes to the buffer.
	w.Close()

	fmt.Println(b.Bytes())
	// Write the compressed data to standard out.
	// os.Stdout implements the io.Writer interface.
	io.Copy(os.Stdout, &b)

}
