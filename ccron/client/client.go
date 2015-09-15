package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ApiError struct {
	Error string `json:"error"`
}

func Request(verb string, server string, path string, reqBody io.Reader) ([]byte, error) {
	c := &http.Client{}
	req, err := http.NewRequest(verb, fmt.Sprintf("%s%s", server, path), reqBody)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 400 {
		data := &ApiError{}
		err := json.Unmarshal(respBody, &data)
		if err == nil {
			return nil, fmt.Errorf(data.Error)
		}

		return nil, fmt.Errorf("ERROR: %d response from server: %s", resp.StatusCode, respBody)
	}

	return respBody, nil
}
