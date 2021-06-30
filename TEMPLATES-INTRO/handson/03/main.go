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

type restaurant struct {
	Name string
	Menu menu
}

type menu struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

func main() {

	r := restaurant{
		Name: "food house",
		Menu: menu{
			Breakfast: []string{"pancakes", "waffles"},
			Lunch:     []string{"borger", "sandwich"},
			Dinner:    []string{"pastas", "eyes"},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}
}
