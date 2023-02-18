package main

import "fmt"

func main() {
    a := 'a'
    var b rune = 'b'
    c := '🧯'


   fmt.Printf("Value of a is %v and type is %T\n", a, a)
   fmt.Printf("Value of b is %v and type is %T\n", b, b)
   fmt.Printf("Value of c is %v and type is %T\n", c, c)

d := rune(0x1F9EF)
fmt.Printf("Value of d is %v and type is %T\n", d, d)

f := string('a')
fmt.Println(f)

greeting := "Runes are cool!"
fmt.Println(string([]rune(greeting)))

af :="生活很好 🔒"
fmt.Println(string(af[len(af)-1]))

fmt.Printf("%v", len([]rune(af)))
fmt.Println(len('🧯')) // 1

}
