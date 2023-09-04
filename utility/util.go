package util

import "net/http"

func CheckMethod(request *http.Request, writer http.ResponseWriter, method string) {
	if request.Method != method {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
