package main

import (
	"os"
	"log"
	"io"
	"strings"
)

func main() {
	name := "Richard Tseng"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	// create a file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	// write template to the new file
	io.Copy(nf, strings.NewReader(tpl))
}