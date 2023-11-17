package apirequest

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get(endpoint string, response interface{}) error {

	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	//We Read the response body on the line below.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	tmp := ApiResponseSuccess{}
	if err = json.Unmarshal(data, &tmp); err == nil {
		if err = tmp.GetError(); err != nil {
			return err
		}
	}

	err = nil
	if response != nil {
		err = json.Unmarshal(data, &response)
	}

	return err

}
