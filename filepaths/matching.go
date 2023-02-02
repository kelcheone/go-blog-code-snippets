package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    files, _ := filepath.Glob("*.txt")
    fmt.Println(files)
}
