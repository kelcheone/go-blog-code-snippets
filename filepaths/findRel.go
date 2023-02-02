package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    fmt.Println(filepath.Rel("home/use.rs", "home/user/file.txt"))
}
