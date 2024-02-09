package functions

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

func ProcessText(s string) (string, int, bool) {
	for _, ch := range s {
		if (ch < 32 && ch != 10) || ch > 126 {
			return "", -1, false
		}
	}

	res := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			count++
			res += string(s[i])
		} else {
			res += string(s[i])
		}
	}
	return res, count, true
}

func ReadWholeFile(bannerChoice string) (map[int]string, error) {
	var filename string
	switch bannerChoice {
	case "standard":
		filename = "banners/standard.txt"
	case "shadow":
		filename = "banners/shadow.txt"
	case "thinkertoy":
		filename = "banners/thinkertoy.txt"
	default:
		return nil, errors.New("Invalid banner")
	}

	sizesMap := map[string]int64{
		"banners/standard.txt":   6623,
		"banners/shadow.txt":     7463,
		"banners/thinkertoy.txt": 5558,
	}

	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("notFound")
	}

	fileInfo, _ := os.Stat(filename)
	if fileInfo.Size() != sizesMap[filename] {
		return nil, errors.New("fileModified")
	}

	text := string(contents)
	return createMap(text, filename), nil
}

func createMap(text, filename string) map[int]string {
	symbolRegex := regexp.MustCompile(`([^\n]+\n){8}`)

	if filename == "banners/thinkertoy.txt" {
		symbolRegex = regexp.MustCompile(`([^\r\n]+\r\n){8}`)
	}

	asciiArtMap := make(map[int]string)
	asciiArtArr := symbolRegex.FindAllString(text, 95)

	for i, symbol := range asciiArtArr {
		symbol = strings.ReplaceAll(symbol, "\r\n", "\n")
		asciiArtMap[i+32] = symbol
	}
	return asciiArtMap
}
