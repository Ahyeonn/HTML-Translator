package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func readHtmlFromFile(fileName string) (string, error) {

	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func parse(text string, tag string) (data []string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string

	var isLi bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			isLi = t.Data == tag

		case tt == html.TextToken:

			t := tkn.Token()

			if isLi {
				vals = append(vals, t.Data)
			}

			isLi = false
		}
	}
}

type Page struct {
	TextFilePath   string
	TranslatedPage string
	Title          string
	Content        string
}

func main() {
	convert := flag.String("convert", "", "convert html to preferred language")
	flag.Parse()

	// call this flag to indicate which html to translate
	if *convert != "" {
		fileName := (*convert)
		text, err := readHtmlFromFile(fileName)

		if err != nil {
			log.Fatal(err)
		}

		fileTitle := parse(text, "h1")[0] // exists only one
		fileContent := parse(text, "p")[0]

		pageName := string("translate.html")
		// fmt.Println(title[0], content[0])
		page := Page{
			TextFilePath:   "./",
			TranslatedPage: pageName,
			Title:          string(fileTitle),
			Content:        string(fileContent),
		}
		t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
		newFile, err := os.Create(page.TranslatedPage)
		if err != nil {
			panic(err)
		}
		t.Execute(newFile, page)
		fmt.Print(">> New translated page has been created")
		// }
	}
}
