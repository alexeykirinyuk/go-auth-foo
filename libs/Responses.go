package libs

import "net/http"

func BadRequest(writer http.ResponseWriter, request *http.Request, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		panic("error when writing to response")
	}
	request.Response.StatusCode = 400
}

func NotAuthorized(writer http.ResponseWriter, request *http.Request, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		panic("error when writing to response")
	}
	request.Response.StatusCode = 401
}
