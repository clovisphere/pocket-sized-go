package main

import (
	"flag"
	"fmt"
)

// language represents the language's code.
type language string

// phrasebook holds greeting for each supported language.
var phraseBook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// main starts program execution.
func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, fr...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// greet says hello to the world in a specified language.
func greet(l language) string {
	greeting, ok := phraseBook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
