package apirequest

import (
	"encoding/json"
	"io"
	"net/http"
)

func Delete(endpoint string) error {
	client := &http.Client{}

	// create a new DELETE request
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	apiresp := ApiResponseSuccess{}
	err = json.Unmarshal(data, &apiresp)
	if err != nil {
		return err
	}

	return apiresp.GetError()
}
