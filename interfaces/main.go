package main

type Printer interface {
	printIt()
}

type MyPrinter struct {
}

func (p *MyPrinter) printIt() {
	println("printIt")
}

func main() {
	var p Printer = &MyPrinter{}
	p.printIt()
}
