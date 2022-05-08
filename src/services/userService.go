package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const endpoint = "http://172.17.0.1:8001/api/"

func Request(method string, path string, cookie string, body map[string]string) (*http.Response, error) {
	var data io.Reader = nil

	if body != nil {
		jsonData, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		data = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, endpoint+path, data)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if cookie != "" {
		req.Header.Add("Cookie", "jwt="+cookie)
	}

	client := &http.Client{}

	return client.Do(req)
}
