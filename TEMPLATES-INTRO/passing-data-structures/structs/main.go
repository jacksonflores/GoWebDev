package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type homie struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	jackson := homie{
		Name: "jackson",
		Age:  18,
	}

	err := tpl.Execute(os.Stdout, jackson)
	if err != nil {
		log.Fatalln(err)
	}
}
