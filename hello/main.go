package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang,
		"lang",
		"en",
		"The required langugage, e.g. en, ur...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"el": "Γεια σου κόσμε",    // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào thế giới", // Vietnamese
}

// greet returns a greeting to the world.
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
