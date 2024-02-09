package functions

import (
	"net/http"
	"strings"
	"text/template"
)

type PageVariables struct {
	InputTxt  string
	OutputTxt string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	mainTmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	err = mainTmpl.Execute(w, nil)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	inputTxt := strings.ReplaceAll(r.FormValue("input"), "\r\n", "\n")
	bannerChoice := r.FormValue("banner")
	// fmt.Printf("%#v\n", inputTxt)

	asciiArtMap, err := ReadWholeFile(bannerChoice)
	if err != nil {
		if err.Error() == "notFound" {
			ErrorPageHandler(w, http.StatusNotFound)
			return
		} else if err.Error() == "Invalid banner" {
			ErrorPageHandler(w, http.StatusMethodNotAllowed)
			return
		} else {
			ErrorPageHandler(w, http.StatusInternalServerError)
			return
		}
	}

	res, isValid := GetAscii(inputTxt, asciiArtMap)
	if !isValid {
		ErrorPageHandler(w, http.StatusBadRequest)
		return
	}

	mainTmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	info := PageVariables{inputTxt, res}

	if r.Method != "POST" {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
		return
	}

	err = mainTmpl.Execute(w, info)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}

func ErrorPageHandler(w http.ResponseWriter, statCode int) {
	pages := map[int]string{
		http.StatusBadRequest:          "templates/HTTP400.html",
		http.StatusNotFound:            "templates/HTTP404.html",
		http.StatusInternalServerError: "templates/HTTP500.html",
		http.StatusMethodNotAllowed:    "templates/HTTP405.html",
	}

	w.WriteHeader(statCode)
	template.Must(template.ParseFiles(pages[statCode])).Execute(w, nil)
}

func ServeCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/css/style.css")
	return
}
