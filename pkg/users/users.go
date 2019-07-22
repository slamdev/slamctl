package users

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

type Users interface {
	Create(username string, force bool) error
}

func NewUsers(out io.Writer) Users {
	return &users{out: out}
}

type users struct {
	out io.Writer
}

func (u *users) Create(username string, force bool) error {
	str := fmt.Sprintf("some sample text with inputs: %v %v", username, force)
	_, err := u.out.Write([]byte(str))
	return errors.WithStack(err)
}
