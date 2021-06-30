package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		fname: "jackson",
		lname: "flores",
		age:   17,
	}
	sa1 := secretAgent{
		person: person{
			fname: "james",
			lname: "bond",
			age:   69,
		},
		ltk: true,
	}
	p1.pSpeak()
	sa1.pSpeak()
	sa1.saSpeak()
}

func (p person) pSpeak() {
	fmt.Println(p.fname, p.lname, "is", p.age)
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.fname, sa.lname, "license to kill:", sa.ltk)
}
