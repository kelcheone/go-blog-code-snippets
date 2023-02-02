package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    fmt.Println(filepath.Dir("home/user/file.txt"))
}
