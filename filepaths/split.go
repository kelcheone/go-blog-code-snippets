package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    dir, file := filepath.Split("home/user/file.txt")
    fmt.Println(dir, file)
}
