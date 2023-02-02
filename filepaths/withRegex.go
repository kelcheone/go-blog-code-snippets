package main

import (
    "fmt"
    "path/filepath"
)

func main() {
// get files that have a f or g in the name
    files, _ := filepath.Glob("*[fg]*")
    fmt.Println(files)
}
