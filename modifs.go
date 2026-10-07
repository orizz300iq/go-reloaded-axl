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
			if i > 0 {
				mots[i-1] = strings.ToUpper(mots[i-1])
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
			continue
		case "(low)":
			if i > 0 {
				mots[i-1] = strings.ToLower(mots[i-1])
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
		case "(cap)":
			if i > 0 {
				mots[i-1] = Capitalize(mots[i-1])
			}
			mots = append(mots[:i], mots[i+1:]...)
			i--
		}

	}
}

func FixPonctuation(texte string) string {

}

func FixPhrases(texte string) string {

}

func FixGrammar(texte string) string {

}

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	minus := ""
	for _, lettre := range s {
		if 65 <= lettre && lettre <= 90 {
			minus += string(lettre + 32)
		} else {
			minus += string(lettre)
		}
	}
	nv_mot := ""
	majuscule := true
	for i := range minus {
		if majuscule {
			if 'a' <= minus[i] && minus[i] <= 'z' {
				nv_mot += string(minus[i] - 32)
				majuscule = false
			} else if 'A' <= minus[i] && minus[i] <= 'Z' || '0' <= minus[i] && minus[i] <= '9' {
				nv_mot += string(minus[i])
				majuscule = false
			} else {
				nv_mot += string(minus[i])
			}
		} else if !('a' <= minus[i] && minus[i] <= 'z' || 'A' <= minus[i] && minus[i] <= 'Z' || '0' <= minus[i] && minus[i] <= '9') {
			nv_mot += string(minus[i])
			majuscule = true
		} else {
			nv_mot += string(minus[i])
		}
	}
	return nv_mot
}
