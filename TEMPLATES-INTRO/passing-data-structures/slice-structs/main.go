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
	davis := homie{
		Name: "Davis",
		Age:  17,
	}

	idiots := []homie{jackson, davis}

	err := tpl.Execute(os.Stdout, idiots)
	if err != nil {
		log.Fatalln(err)
	}
}
