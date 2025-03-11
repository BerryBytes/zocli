package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/berrybytes/zocli/api"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(req *http.Request) (*http.Response, error) {
	var newUser loginReq

	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}

	if newUser.Password == "" {
		return &http.Response{StatusCode: 400}, errors.New("required Password")
	}
	if !strings.Contains(newUser.Email, "@") {
		return &http.Response{StatusCode: 400}, errors.New("invalid email")
	}

	return successResponse()
}

func TokenLogin(req *http.Request) (*http.Response, error) {
	token := req.Header.Get("X-PERSONAL-TOKEN")

	if token == "" {
		return &http.Response{StatusCode: 400}, errors.New("required Token")
	}

	if token == "validToken" {
		return ssoSuccessResponse()
	}
	return &http.Response{StatusCode: 404}, errors.New("no such token")
}

func BrowserLogin(req *http.Request) (*http.Response, error) {
	switch {
	case req.Method == http.MethodGet:
		return &http.Response{StatusCode: http.StatusMethodNotAllowed}, errors.New("no such method allowed")
	case req.Method == http.MethodPost:
		resToken := []byte(`{"success": 1, "message": "Success", "data": {"code": "abcdef"}}`)
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(resToken))}, nil
	}

	return nil, nil
}

func SSOStatus(req *http.Request) (*http.Response, error) {
	switch {
	case req.Method == http.MethodGet:
		type SSOReq struct {
			SSOToken string `json:"code"`
		}
		var newReq SSOReq
		err := json.NewDecoder(req.Body).Decode(&newReq)
		if err != nil {
			return &http.Response{StatusCode: 500}, errors.New("internal server error")
		}

		fmt.Println(newReq.SSOToken)
		if newReq.SSOToken != "abcdef" {
			return &http.Response{StatusCode: 400}, errors.New("no such token " + newReq.SSOToken)
		}

		return ssoSuccessResponse()
	case req.Method == http.MethodPost:
		return &http.Response{StatusCode: http.StatusMethodNotAllowed}, errors.New("no such method allowed")
	}
	return nil, nil
}

func successResponse() (*http.Response, error) {
	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data: api.Data{
			AuthToken: "newTokenGenerated" + strconv.Itoa(rand.Int()),
			User: api.User{
				Id:    rand.Int(),
				Email: "testmail@mail.com",
			},
		},
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}
func ssoSuccessResponse() (*http.Response, error) {
	send := api.BaseResponse{
		Success: 1,
		Message: "Success",
		Data: api.User{
			Id:    rand.Int(),
			Email: "testmail@mail.com",
		},
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func Profile(req *http.Request) (*http.Response, error) {
	response := []byte(`
{
    "success": 1,
    "message": "Success",
    "data": {
        "email": "testmail@mail.com"
    }
}
`)

	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(response))}, nil
}
