package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const DEFAULT_PORT = "80"
const CODE_HEADER = "X-Code"
const FORMAT_HEADER = "X-Format"
const ORIGIN_HEADER = "X-Original-URI"

type ErrorTemplateData struct {
	Status      string
	StatusText  string
	OriginalURI string
	Format      string
}

func main() {
	port := os.Getenv("HTTP_SERVE_PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	fs := http.FileServer(http.Dir("www/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.HandlerFunc(codeHandler))

	log.Printf("Listening to port %s...", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
	numCode, _ := strconv.ParseInt(r.Header.Get(CODE_HEADER), 10, 32)
	td := ErrorTemplateData{
		Status:      r.Header.Get(CODE_HEADER),
		StatusText:  http.StatusText(int(numCode)),
		OriginalURI: r.Header.Get(ORIGIN_HEADER),
		Format:      GetExtensionForMime(r.Header.Get(FORMAT_HEADER)),
	}

	templatePath := GetTemplateFile(td.Status, td.Format)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error while accessing template: %s", templatePath)
	}

	tmpl.Execute(w, td)
}

func GetExtensionForMime(mime string) (format string) {
	if strings.Contains(mime, "text/html") {
		format = "html"
	} else if strings.Contains(mime, "application/json") {
		format = "json"
	} else {
		format = "html"
	}
	return
}

func GetTemplateFile(code string, format string) string {
	if code == "" {
		code = "404"
	}

	if format == "" {
		format = "html"
	}

	c := []byte(code)
	filePath := fmt.Sprintf("www/%s.%s", code, format)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		for i := (len(c) - 1); i >= 0; i-- {
			if i == (len(c)-1) && c[i] != 'x' || c[i] != 'x' {
				c[i] = 'x'
				break
			}
		}

		if string(c) == "xxx" {
			c = []byte("404")
		}

		return GetTemplateFile(string(c), format)
	}
	return filePath
}
