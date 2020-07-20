package routes

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func WriteContent(w http.ResponseWriter, r *http.Request, templatePath string) {
	content, err := read(templatePath)
	if err != nil {
		fmt.Printf("error when read template '%s': %s", templatePath, err)
		InternalServerError(w, r)
		return
	}

	_, err = w.Write([]byte(content))
	if err != nil {
		fmt.Printf("error when write response '%s': %s", templatePath, err)
		InternalServerError(w, r)
		return
	}
}

func Render(w http.ResponseWriter, r *http.Request, templatePath string, data interface{}) {
	content, err := read(templatePath)
	if err != nil {
		fmt.Printf("error when read template '%s': %s", templatePath, err)
		InternalServerError(w, r)
		return
	}

	tpl, err := template.New(templatePath).Parse(content)
	if err != nil {
		fmt.Printf("error when parse template '%s': %s", templatePath, err)
		InternalServerError(w, r)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Printf("error when generate view from template '%s': %s", templatePath, err)
		InternalServerError(w, r)
		return
	}
}

func read(filePath string) (content string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			// TODO
			panic("")
		}
	}()

	byteArray, err := ioutil.ReadAll(file)
	content = string(byteArray)
	return
}
