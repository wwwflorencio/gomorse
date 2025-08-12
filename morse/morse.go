package morse

import (
	"strings"
)

func getAlphabet() map[string]string {
	return map[string]string{
		".-":    "A",
		"-...":  "B",
		"-.-.":  "C",
		"-..":   "D",
		".":     "E",
		"..-.":  "F",
		"--.":   "G",
		"....":  "H",
		"..":    "I",
		".---":  "J",
		"-.-":   "K",
		".-..":  "L",
		"--":    "M",
		"-.":    "N",
		"---":   "O",
		".--.":  "P",
		"--.-":  "Q",
		".-.":   "R",
		"...":   "S",
		"-":     "T",
		"..-":   "U",
		"...-":  "V",
		".--":   "W",
		"-..-":  "X",
		"-.--":  "Y",
		"--..":  "Z",
		"-----": "0",
		".----": "1",
		"..---": "2",
		"...--": "3",
		"....-": "4",
		".....": "5",
		"-....": "6",
		"--...": "7",
		"---..": "8",
		"----.": "9",
	}
}

type MorseDecoder struct {
	letterSeparator       string
	wordSeparator         string
	invalidLetterFallback string
	alphabet              map[string]string
}

func NewMorseDecoder() *MorseDecoder {
	return &MorseDecoder{
		letterSeparator:       " ",
		wordSeparator:         "   ",
		alphabet:              getAlphabet(),
		invalidLetterFallback: "?",
	}
}

func (md *MorseDecoder) Decode(code string) string {
	code = strings.TrimSpace(code)
	
	if code == "" {
		return ""
	}

	words := strings.Split(code, md.wordSeparator)

	var decodedWords []string
	for _, word := range words {
		letters := strings.Split(word, md.letterSeparator)

		var builder strings.Builder

		for _, letter := range letters {
			if val, exists := md.alphabet[letter]; exists {
				builder.WriteString(val)
			} else {
				builder.WriteString(md.invalidLetterFallback)
			}
		}

		decodedWords = append(decodedWords, builder.String())
	}

	return strings.Join(decodedWords, " ")
}
