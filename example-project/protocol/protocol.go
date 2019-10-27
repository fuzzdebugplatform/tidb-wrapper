package protocol

import (
	"example/server"
	"net/http"
)

func Handle() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query()["key"]
		server.ProcessKey(key[0], writer)
	}
}
