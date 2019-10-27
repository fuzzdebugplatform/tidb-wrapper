package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func ProcessKey(key string, writer http.ResponseWriter) {
	keyInt, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("Error strconv atoi")
	}

	if keyInt > 10 {
		keyInt--
	} else {
		keyInt++
	}

	writer.Write([]byte(strconv.Itoa(keyInt)))
}
