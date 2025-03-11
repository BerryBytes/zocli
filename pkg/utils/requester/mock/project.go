package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/berrybytes/zocli/api"
)

func GetAllProjects(_ *http.Request) (*http.Response, error) {
	response := []byte(`
{
	"success": 1,
    "message": "Success",
	"data": [
        {
            "id": 529,
            "name": "Hem Test",
            "project_code": "hello",
            "tags": "new tag,old tag",
            "region": "",
            "active": true,
            "cluster_scope": 0,
            "subscription": {
                "id": 2,
                "createdat": "2020-08-22T07:22:52.45931Z",
                "name": "Business Plan",
                "resource_list": {
                    "pods": "100",
                    "secrets": "100",
                    "services": "100"
                },
                "load_balancer": 5,
                "price_list": null
            }
		}
	]
}`)
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(response))}, nil
}

func GetSingleProject(req *http.Request) (*http.Response, error) {
	ids := strings.Split(req.URL.String(), "/")
	id := ids[len(ids)-1]

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
	var project api.Project
	project.Name = "Mocked Project"
	project.Description = "mocked Description"
	project.ID = 1

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    project,
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func EnableProjectByID(req *http.Request) (*http.Response, error) {
	ids := strings.Split(req.URL.String(), "/")
	id := ids[len(ids)-2]

	if id != "1" {
		return &http.Response{StatusCode: 400}, errors.New("no such id found")
	}
	response := []byte(`
{
	"success": 1,
    "message": "Success",
	"data": null
	}`)
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewBuffer(response))}, nil
}

func GetProjectByName(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	name := url[len(url)-1]

	if name != "testProj" {
		return &http.Response{StatusCode: 400}, errors.New("no such project found")
	}

	var project api.Project
	project.Name = "Mocked Project"
	project.Description = "mocked Description"
	project.ID = 1

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    project,
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func DisableProjectByID(req *http.Request) (*http.Response, error) {
	ids := strings.Split(req.URL.String(), "/")
	id := ids[len(ids)-2]

	if id != "1" {
		return &http.Response{StatusCode: 400}, errors.New("no such id found")
	}
	response := []byte(`
{
	"success": 1,
    "message": "Success",
	"data": null
	}`)
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewBuffer(response))}, nil
}

func RenameProject(req *http.Request) (*http.Response, error) {
	type Req struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	}

	var rename Req
	err := json.NewDecoder(req.Body).Decode(&rename)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	ids := strings.Split(req.URL.String(), "/")
	id, _ := strconv.Atoi(ids[len(ids)-1])
	rename.ID = int64(id)

	if rename.ID != 1 {
		return &http.Response{StatusCode: 400}, errors.New("no such id")
	}

	var project api.Project
	project.Name = rename.Name
	project.ID = 1

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    project,
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func GetSingleProjectVars(req *http.Request) (*http.Response, error) {
	ids := strings.Split(req.URL.String(), "/")
	id := ids[len(ids)-1]

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
	var project api.Project
	project.Name = "Mocked Project"
	project.Description = "mocked Description"
	project.ID = 1
	project.Variables = []api.Variable{
		{Id: 1, Key: "test", Type: "secret", Value: "test"},
	}

	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    project,
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func DeleteVariable(req *http.Request) (*http.Response, error) {
	a := struct {
		Variables []api.Variable `json:"variables"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&a)
	if err != nil {
		return nil, err
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

func ListProjectPermissions(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	projectID := url[len(url)-2]

	if projectID != "1" {
		fmt.Println("the project id doesn't match")
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

	var list api.Permissions
	list.Permissions = append(list.Permissions, api.Permission{
		Id:        1,
		CreatedAt: time.Now().Add(-100 * time.Minute),
		Email:     "test@mail.com",
		User: api.User{
			Id:            2,
			FirstName:     "Test",
			LastName:      "Test",
			Email:         "test@mail.com",
			Company:       "test",
			EmailVerified: true,
		},
		UserId: 2,
		UserRole: api.UserRole{
			Id:          3,
			CreatedAt:   time.Now(),
			Name:        "Read",
			Code:        3,
			Description: "Analyst role",
			Active:      true,
		},
		UserRoleId: 3,
		Active:     true,
	})
	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    list.Permissions,
	}
	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func GetAllRoles(req *http.Request) (*http.Response, error) {
	var roles api.Roles
	roles.Roles = append(roles.Roles, api.Role{
		Id:          1,
		CreatedAt:   time.Now(),
		Name:        "Admin",
		Code:        1,
		Description: "this is admin description",
		Active:      true,
	})
	send := api.BaseResponse{
		Success: 1,
		Message: "success",
		Data:    roles,
	}

	r, err := json.Marshal(send)
	if err != nil {
		return &http.Response{StatusCode: 500}, errors.New("internal server error")
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(r))}, nil
}

func DeletePermission(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	id := url[len(url)-1]
	if id != "1" {
		fmt.Println("the project id doesn't match")
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

func UpdatePermissions(req *http.Request) (*http.Response, error) {
	url := strings.Split(req.URL.String(), "/")
	id := url[len(url)-1]
	if id != "1" {
		fmt.Println("the project id doesn't match")
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
