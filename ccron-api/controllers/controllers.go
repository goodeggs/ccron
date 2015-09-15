package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RenderError(rw http.ResponseWriter, err error) error {
	if err != nil {
		http.Error(rw, fmt.Sprintf(`{"error":%q}`, err.Error()), http.StatusInternalServerError)
	}

	return err
}

func RenderSuccess(rw http.ResponseWriter) error {
	_, err := rw.Write([]byte(`{"success":true}`))

	return err
}

func RenderText(rw http.ResponseWriter, text string) error {
	_, err := rw.Write([]byte(text))
	return err
}

func RenderJson(rw http.ResponseWriter, object interface{}) error {
	data, err := json.MarshalIndent(object, "", "  ")

	if err != nil {
		return RenderError(rw, err)
	}

	data = append(data, '\n')

	rw.Header().Set("Content-Type", "application/json")

	_, err = rw.Write(data)

	return err
}
