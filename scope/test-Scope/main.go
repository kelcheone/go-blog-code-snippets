package main

import (
    "fmt"
    "scope-tut/store"
)

func main(){
    fmt.Printf("We are accessing someVar with value %v.", store.SomeVar)
    store.DataIsBeautiful()
}
