package main

import (
	"strconv"
	"strings"
)

func ProcessText (input string) string {
	
	words := strings.Fields(input)

	words = ProcessWords(words)

	text := strings.Join(words, " ")

	text = FixPonctuation(text)
	
	text = FixPhrases(text)

	text = FixGrammar(text)

	return text
}

func ProcessWords(mots []string) []string {
	for i := 0; i < len(mots); i++ {
		switch mots[i] {
			
		case "(hex)":
			if i > 0 {
				val, err := strconv.ParseInt(mots[i-1], 16, 64)
				if err == nil {
					mots[i-1] = strconv.FormatInt(val, 10)
				}
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
		
		case "(bin)":
			if i > 0 {
				val, err := strconv.ParseInt(mots[i-1], 2, 64)
				if err == nil {
					mots[i-1] = strconv.FormatInt(val, 10)
				}
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
		
		case "(up)":

		case "(low)":
			
		case "(cap)":
			
		}

	}
}

func FixPonctuation(texte string) string {

}

func FixPhrases(texte string) string {

}

func FixGrammar(texte string) string {

}