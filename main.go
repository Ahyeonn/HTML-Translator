package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func main() {
	trans := flag.String("trans", "", "google translate")
	flag.Parse()

	if *trans != "" {
		// Set up a context and authenticate with Google Cloud
		ctx := context.Background()
		creds := option.WithCredentialsFile("./credential.json")
		client, err := translate.NewClient(ctx, creds)
		if err != nil {
			panic(err)
		}
		lenArg := os.Args[1][8:]
		// Translate a string
		text := lenArg
		targetLang := language.Make("es")
		resp, err := client.Translate(ctx, []string{text}, targetLang, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp[0].Text)
	}
}
