package main

import "fmt"

func main(){
    num := 10
    if num > 5 {
        defer fmt.Println("num is greater than 5")
    }else{
        defer fmt.Println("num is less than 5")
    }

    fmt.Println("num is", num)
}
