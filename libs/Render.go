package libs

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func Render(w http.ResponseWriter, r *http.Request, templatePath string, data interface{}) {
	content, err := read(templatePath)
	if err != nil {
		panic(err)
	}

	tpl, err := template.New(templatePath).Parse(content)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
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
