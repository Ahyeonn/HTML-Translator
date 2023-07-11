package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/net/html"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
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
		fmt.Print("\n Spanish: es \n French: fr \n Portuguese: pt \n English: en \n Russian: ru \n German: de \n Arabic: ar \n Chinese: zh \n Hindi: hi")

		var languageChoice string
		fmt.Print("\nPlease type the language of the choice : ")
		fmt.Scanln(&languageChoice)

		fileName := (*convert)
		text, err := readHtmlFromFile(fileName)

		if err != nil {
			log.Fatal(err)
		}

		fileTitle := parse(text, "h1")[0]
		fileContent := parse(text, "p")[0]

		ctx := context.Background()
		creds := option.WithCredentialsFile("./credential.json")
		client, err := translate.NewClient(ctx, creds)
		if err != nil {
			panic(err)
		}
		transTitle := fileTitle
		transContent := fileContent
		// Translate a string
		textTitle := transTitle
		textContent := transContent
		targetLang := language.Make(languageChoice)
		resp, err := client.Translate(ctx, []string{textTitle, textContent}, targetLang, nil)
		titleString := resp[0].Text
		titleContent := resp[1].Text
		if err != nil {
			panic(err)
		}

		pageName := string("translate.html")
		page := Page{
			TextFilePath:   "./",
			TranslatedPage: pageName,
			Title:          titleString,
			Content:        titleContent,
		}
		t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
		newFile, err := os.Create(page.TranslatedPage)
		if err != nil {
			panic(err)
		}
		t.Execute(newFile, page)
		fmt.Print(">> New translated page has been created")
	}
}
