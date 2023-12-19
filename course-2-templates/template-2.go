package main 

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	// executing os.Stdout (writer interface)
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}