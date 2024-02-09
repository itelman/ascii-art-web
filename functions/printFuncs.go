package functions

import (
	"regexp"
	"strings"
)

func ConvertToArr(s string) []string {
	state := regexp.MustCompile(`([^\n]+)\n`)
	res := []string{}
	for state.MatchString(s) {
		text := state.FindString(s)
		text = state.ReplaceAllString(text, "$1")
		res = append(res, text)
		indexes := state.FindStringIndex(s)
		s = s[indexes[1]:]
	}
	return res
}

func Join(asciiArtMap map[int]string, text []string, count int) string {
	res := ""
	for _, word := range text {
		arr := [][]string{}
		if len(word) != 0 {
			for _, ch := range word {
				arr = append(arr, ConvertToArr(asciiArtMap[int(ch)]))
			}
			for r := 0; r < 8; r++ {
				for i := 0; i < len(word); i++ {
					res += arr[i][r]
				}
				res += "\n"
			}
		} else if word == "" && count > 0 {
			res += "\n"
			count--
		}
	}
	return res
}

func GetAscii(inputText string, asciiArtMap map[int]string) (string, bool) {
	txt, numSlashN, isValid := ProcessText(inputText)
	if !isValid {
		return "", false
	}
	text := strings.Split(txt, "\n")
	res := Join(asciiArtMap, text, numSlashN)
	return res, true
}
