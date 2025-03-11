package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/berrybytes/zocli/api"
)

func GetAllEnvironment(req *http.Request) (*http.Response, error) {
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

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data: []api.Environment{
			{ID: 1},
		},
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func GetSingleEnvironment(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	fmt.Printf("url: %v\n", url)
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
		Data: api.Environment{
			ID: 1,
		},
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func GetEnvironmentOverview(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	fmt.Printf("url: %v\n", url)
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

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    api.EnvironmentOverview{},
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func StopEnvironment(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	fmt.Printf("url: %v\n", url)
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

	send := api.BaseResponse{
		Success: 1,
		Message: "Sent command for stopping environment",
	}

	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func StartEnvironment(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	fmt.Printf("url: %v\n", url)
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

	send := api.BaseResponse{
		Success: 1,
		Message: "Sent command for starting environment",
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}

	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func DeleteEnvironment(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	fmt.Printf("url: %v\n", url)
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
		Message: "Success",
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}

	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func RenameEnvironment(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	fmt.Printf("url: %v\n", url)
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
		Message: "Success",
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}

	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}
