package internal

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"slamctl/pkg/users"
)

func init() {
	usersRoute := apiRoute.Subrouter().Name("users").PathPrefix("/users")
	usersRoute.Subrouter().Name("users-create").Path("/create").Methods("POST").HandlerFunc(createUser)
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var b createUserBody
	err := decoder.Decode(&b)
	if err != nil {
		sendError(writer, 400, err)
		return
	}

	u := users.NewUsers(writer)
	err = u.Create(b.Username, b.Force)
	if err != nil {
		sendError(writer, 400, err)
	}
}

func sendError(writer http.ResponseWriter, statusCode int, err error) {
	logrus.WithField("error", fmt.Sprintf("%+v", err)).Error()
	writer.WriteHeader(statusCode)
	_, writeError := writer.Write([]byte(err.Error()))
	if writeError != nil {
		writeError = errors.Wrap(writeError, "failed to write response")
		logrus.WithField("error", fmt.Sprintf("%+v", writeError)).Error()
	}
}

type createUserBody struct {
	Username string `json:"username"`
	Force    bool   `json:"force"`
}
