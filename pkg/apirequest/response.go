package apirequest

import "errors"

type ApiResponseSuccess struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *ApiResponseSuccess) GetError() error {
	if r.Success == false {
		return errors.New(r.Message)
	}
	return nil
}
