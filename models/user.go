package models

import (
	"log"
	"net/http"
	"regexp"
)
//easyjson:json
type User struct {
	ID       int64  `json:"-"`
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}
//easyjson:json
type Users []*User

//easyjson:json
type UserParams struct {
	Limit int
	Since string
	Desc  bool
}
//easyjson:json
type UpdateUserFields struct {
	About    *string `json:"about"`
	Email    *string `json:"email"`
	Fullname *string `json:"fullname"`
}

var (
	nicknameRegexp *regexp.Regexp
	emailRegexp    *regexp.Regexp
)

func init() {
	var err error
	nicknameRegexp, err = regexp.Compile(`^[a-zA-Z0-9_.]+$`)
	if err != nil {
		log.Fatalf("nickname regexp err: %s", err.Error())
	}

	emailRegexp, err = regexp.Compile("^[a-zA-Z0-9.!#$%&''*+/=?^_`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$`)
	if err != nil {
		log.Fatalf("email regexp err: %s", err.Error())
	}
}

func (u *User) Validate() *Error {
	if !(nicknameRegexp.MatchString(u.Nickname) &&
		emailRegexp.MatchString(u.Email) &&
		u.Fullname != "") {
		return NewError(http.StatusBadRequest, "validation failed")
	}

	return nil
}