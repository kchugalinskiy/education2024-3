package echoservice

import (
	"fmt"
	"io"
	"net/http"
)

type Service struct {
}

// func(w http.ResponseWriter, r *http.Request)

func copyResponse(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}
	w.WriteHeader(http.StatusOK)
	_, err := io.Copy(w, r.Body)
	if err != nil {
		fmt.Println(err)
	}
}

func copyReadAll(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println(err)
	}
}

func (e *Service) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	/*
		[]struct{Username string}{Username:"vasya"} -> [{"username":"vasya"}] // encoding/json
		GET /users -> `[{"username":"vasya"},{"username":"petya"}]` -> 200
		POST /users <- `{"username":"vasya"}` -> 201
		DELETE /users/vasya -> 204 // request.URL `import "regexp"`

		if method == POST { // request.Method == http.MethodPost
			createUser(w,r)
		} else if method == GET {
			getUser(w,r)
		} ....
	*/
}

func NewService() *Service {
	return &Service{}
}
