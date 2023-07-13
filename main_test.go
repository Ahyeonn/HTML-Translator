package main

import (
	"context"
	"testing"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type testCase struct {
	InputFile         string
	ExpectedTitle     string
	ExpectedContent   string
	TranslationLang   string
	ExpectedTitlePL   string
	ExpectedContentPL string
}

var testCases = []testCase{
	// Spanish
	{
		InputFile:         "test.html",
		ExpectedTitle:     "English Title",
		ExpectedContent:   "English Content",
		TranslationLang:   "es",
		ExpectedTitlePL:   "Título de la prueba",
		ExpectedContentPL: "Contenido de prueba",
	},
	// German
	{
		InputFile:         "test.html",
		ExpectedTitle:     "English Title",
		ExpectedContent:   "English Content",
		TranslationLang:   "de",
		ExpectedTitlePL:   "Testtitel",
		ExpectedContentPL: "Testinhalte",
	},
}

func TestMain(t *testing.T) {
	for _, tc := range testCases {
		text, err := readHtmlFromFile(tc.InputFile)
		if err != nil {
			t.Errorf("failed to read HTML file: %v", err)
			continue
		}

		fileTitle := parse(text, "h1")[0]
		fileContent := parse(text, "p")[0]

		ctx := context.Background()
		creds := option.WithCredentialsFile("./credential.json")
		client, err := translate.NewClient(ctx, creds)
		if err != nil {
			t.Errorf("failed to create translation client: %v", err)
			continue
		}

		transTitle := fileTitle
		transContent := fileContent
		textTitle := transTitle
		textContent := transContent
		targetLang := language.Make(tc.TranslationLang)
		resp, err := client.Translate(ctx, []string{textTitle, textContent}, targetLang, nil)
		if err != nil {
			t.Errorf("translation failed: %v", err)
			continue
		}

		titleString := resp[0].Text
		titleContent := resp[1].Text

		if titleString != tc.ExpectedTitlePL {
			t.Errorf("unexpected translated title: got %q, want %q", titleString, tc.ExpectedTitlePL)
		}
		if titleContent != tc.ExpectedContentPL {
			t.Errorf("unexpected translated content: got %q, want %q", titleContent, tc.ExpectedContentPL)
		}
	}
}