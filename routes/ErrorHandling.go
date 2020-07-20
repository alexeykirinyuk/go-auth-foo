package routes

import "net/http"

func InternalServerError(writer http.ResponseWriter, request *http.Request)  {
	_, err := writer.Write([]byte("Internal Server Error"))
	if err != nil {
		panic("error when writing to response")
	}
	request.Response.StatusCode = 500
}
