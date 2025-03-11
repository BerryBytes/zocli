package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/berrybytes/zocli/api"
)

func ListApps(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	id := url[len(url)-2]
	if id != "1" {
		send := api.BaseResponse{
			Success: 0,
			Message: "invalid id",
		}
		r, err := json.Marshal(send)
		if err != nil {
			return &http.Response{StatusCode: 500}, errors.New("internal server error")
		}
		return &http.Response{StatusCode: 400, Body: io.NopCloser(bytes.NewReader(r))}, nil
	}

	appsList := []api.Application{
		{Id: 1, Name: "test"},
	}

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    appsList,
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func RenameApplication(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	id := url[len(url)-1]
	if id != "1" {
		send := api.BaseResponse{
			Success: 0,
			Message: "invalid id",
		}
		r, err := json.Marshal(send)
		if err != nil {
			return &http.Response{StatusCode: 500}, errors.New("internal server error")
		}
		return &http.Response{StatusCode: 400, Body: io.NopCloser(bytes.NewReader(r))}, nil
	}

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func DeleteApplication(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	id := url[len(url)-1]
	if id != "1" {
		send := api.BaseResponse{
			Success: 0,
			Message: "invalid id",
		}
		r, err := json.Marshal(send)
		if err != nil {
			return &http.Response{StatusCode: 500}, errors.New("internal server error")
		}
		return &http.Response{StatusCode: 400, Body: io.NopCloser(bytes.NewReader(r))}, nil
	}

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}
