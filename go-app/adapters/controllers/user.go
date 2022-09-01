package controllers

import (
	"database/sql"
	"net/http"
	"strings"

	"go-app/usecases/user"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) user.UserOutputPort
	InputFactory  func(o user.UserOutputPort, u user.UserRepository) user.UserInputPort
	RepoFactory   func(c *sql.DB) user.UserRepository
	Conn          *sql.DB
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	outputPort := u.OutputFactory(w)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(ctx, userID)
}
