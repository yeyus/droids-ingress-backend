package main

import "testing"

type getTemplateFileTestCase struct {
	code   string
	format string
	path   string
}

func TestGetTemplateFile(t *testing.T) {

	var testCases = []getTemplateFileTestCase{
		{code: "700", format: "html", path: "www/404.html"},
		{code: "404", format: "html", path: "www/404.html"},
		{code: "404", format: "json", path: "www/404.json"},
		{code: "200", format: "json", path: "www/404.json"},
		{code: "429", format: "json", path: "www/4xx.json"},
		{code: "500", format: "html", path: "www/500.html"},
		{code: "500", format: "json", path: "www/500.json"},
		{code: "503", format: "html", path: "www/5xx.html"},
		{code: "503", format: "json", path: "www/5xx.json"},
		{code: "", format: "html", path: "www/404.html"},
		{code: "", format: "json", path: "www/404.json"},
		{code: "404", format: "", path: "www/404.html"},
		{code: "", format: "", path: "www/404.html"},
	}

	var v string
	for i, tc := range testCases {
		v = GetTemplateFile(tc.code, tc.format)
		if v != tc.path {
			t.Errorf("test case %d -> Invalid output for code: %s and format: %s, expected %s, got %s", i, tc.code, tc.format, tc.path, v)
		}
	}
}

type getExtensionForMimeTestCase struct {
	mime   string
	result string
}

func TestGetExtensionForMime(t *testing.T) {

	var testCases = []getExtensionForMimeTestCase{
		{mime: "text/html", result: "html"},
		{mime: "text/plain", result: "plain"},
		{mime: "application/json", result: "json"},
		{mime: "text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8", result: "html"},
		{mime: "text/html;charset=UTF-8", result: "html"},
	}

	var v string
	for i, tc := range testCases {
		v = GetExtensionForMime(tc.mime)
		if v != tc.result {
			t.Errorf("test case %d -> Invalid output for mime: %s, expected %s, got %s", i, tc.mime, tc.result, v)
		}
	}
}
