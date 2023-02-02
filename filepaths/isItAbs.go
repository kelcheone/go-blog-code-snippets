package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    fmt.Println(filepath.IsAbs("file.txt"))
}
