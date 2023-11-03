package apirequest

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/eysteinn/stackmap-cli/pkg/global"
)

func PostForm(endpoint string, form map[string]string, response interface{}) error {
	ct, form_body, err := createForm(form)
	if err != nil {
		return err
	}
	resp, err := http.Post(endpoint, ct, form_body)
	if err != nil {
		return err
	}
	//We Read the response body on the line below.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Println(string(data))

	tmp := []ApiResponseSuccess{}
	if err = json.Unmarshal(data, &tmp); err == nil {
		for _, r := range tmp {
			if err = r.GetError(); err != nil {
				return err
			}
		}
	}

	err = nil
	if response != nil {
		err = json.Unmarshal(data, response)
	}

	return err
}

func createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)

	mp := multipart.NewWriter(body)
	defer mp.Close()
	for key, val := range form {
		if strings.HasPrefix(val, "@") {
			val = val[1:]
			file, err := os.Open(val)
			if err != nil {
				return "", nil, err
			}
			defer file.Close()
			part, err := mp.CreateFormFile(key, val)
			if err != nil {
				return "", nil, err
			}
			io.Copy(part, file)
		} else {
			mp.WriteField(key, val)
		}
	}

	if !global.Quiet {
		bar := pb.New(body.Len()).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
		bar.ShowSpeed = true
		bar.Start()
		reader := bar.NewProxyReader(body)
		return mp.FormDataContentType(), reader, nil
	}
	return mp.FormDataContentType(), body, nil
}
