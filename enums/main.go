package main

import "fmt"

type Subject int

const (
	English Subject = iota + 1
	Kiswahili
	Mathematics
	Physics
	Chemistry
	Biology
)

// String returns the string representation of the Subject
func (s Subject) String() string {
	return [...]string{"English", "Kiswahili", "Mathematics", "Physics", "Chemistry", "Biology"}[s-1]
}

// Index returns the index of the Subject
func (s Subject) Index() int {
	return int(s)
}

func main() {
	subject := Biology
	fmt.Println(subject)          // Biology
	fmt.Println(subject.Index())  // 6
	fmt.Println(subject.String()) // Biology
}
