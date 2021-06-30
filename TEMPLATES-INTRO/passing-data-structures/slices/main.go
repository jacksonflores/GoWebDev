package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	homies := []string{"davis", "charlie", "scott"}

	err := tpl.Execute(os.Stdout, homies)
	if err != nil {
		log.Fatalln(err)
	}
}
