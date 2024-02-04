package tools

import (
	"io"
	"log"
	"net/http"
)

func ReadResponse(resp *http.Response) string {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}
