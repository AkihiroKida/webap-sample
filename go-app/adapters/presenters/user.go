package presenters

import (
	"fmt"
	"net/http"

	"go-app/entities"
	"go-app/usecases/user"
)

type User struct {
	w http.ResponseWriter
}

func NewUserOutputPort(w http.ResponseWriter) user.UserOutputPort {
	return &User{
		w: w,
	}
}

func (u *User) Render(user *entities.User) {
	u.w.WriteHeader(http.StatusOK)
	fmt.Fprint(u.w, user.Name)
}

func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}
