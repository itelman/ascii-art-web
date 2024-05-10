package functions

import (
	"bytes"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type PageVariables struct {
	InputTxt, OutputTxt string
}

type Error struct {
	Code    int
	Name    string
	Content string
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

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art/output.txt" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
		return
	}

	res := r.FormValue("res")

	if len(res) == 0 {
		ErrorPageHandler(w, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=output.txt")
	w.Header().Set("Content-Length", strconv.Itoa(len(res)))
	w.Header().Set("Content-Type", "text/plain")

	http.ServeContent(w, r, "output.txt", time.Now(), bytes.NewReader([]byte(res)))
	return
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
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
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	info := PageVariables{inputTxt, res}

	mainTmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	err = mainTmpl.Execute(w, info)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}

func ErrorPageHandler(w http.ResponseWriter, statCode int) {
	errors := map[int]Error{
		http.StatusBadRequest:          {http.StatusBadRequest, "Bad Request", "The server cannot process the request due to something that is perceived to be a client error."},
		http.StatusNotFound:            {http.StatusNotFound, "Resource not found", "The requested resource could not be found but may be available again in the future."},
		http.StatusMethodNotAllowed:    {http.StatusMethodNotAllowed, "Method Not Allowed", "A request method is not supported for the requested resource."},
		http.StatusInternalServerError: {http.StatusInternalServerError, "Webservice currently unavailable", "An unexpected condition was encountered.<br />Our service team has been dispatched to bring it back online."},
	}

	errTmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(statCode)
		errTmpl.Execute(w, errors[statCode])
	}
}

func ServeCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/css/style.css")
	return
}
