package server

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Output HTTP response. TODO protobuf
func outputHttp(w http.ResponseWriter, data string) {
	io.WriteString(w, data)
}

// Send Http request. TODO protobuf
func sendHttpRemoteNode(url string) (response string, err error) {
	// TODO not return 404 and some error
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
