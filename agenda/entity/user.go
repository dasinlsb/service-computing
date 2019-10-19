package entity

import (
	"agenda/model/user"
	"errors"
	"fmt"
	"strings"
)

func GetCustom(username string) ([][]string, error) {
	session := user.GetSession()
	var info [][]string
	if !session.LoggedIn {
		return info, notloginerr()
	}
	u, find := user.FindOneByName(username)
	if !find {
		return info, errors.New("not found such user")
	}
	info = append(info, []string{"Username", "Email", "Telephone"})
	info = append(info, []string{u.Username, u.Email, u.Phone})
	return info, nil
}

func List() (info [][]string, err error) {
	session := user.GetSession()
	if !session.LoggedIn {
		return info, notloginerr()
	}
	users := user.GetUsers()
	info = append(info, []string{"Username", "Email", "Telephone"})
	for _, u := range users {
		info = append(info, []string{u.Username, u.Email, u.Phone})
	}
	return info, nil
}

func Login(username, password string) error {
	session := user.GetSession()
	if session.LoggedIn {
		return errors.New("you have already logged in")
	}
	return user.Auth(username, password)
}

func Logout() error {
	session := user.GetSession()
	if !session.LoggedIn {
		return notloginerr()
	}
	return user.Logout()
}

func Register(u user.User) error {
	if sects := user.ValidateNewUser(u); len(sects) > 0 {
		return fmt.Errorf(strings.Join(sects, ", ") + " repeated!")
	}
	return user.Add(u)
}

func Status() (info [][]string, err error) {
	session := user.GetSession()
	if !session.LoggedIn {
		info = append(info, []string{"LoggedIn", "NO"})
		return info, notloginerr()
	}
	info = append(info, []string{"LoggedIn", "YES"})
	info = append(info, []string{"Username", session.Username})
	return info, nil
}

func Remove() error {
	session := user.GetSession()
	if !session.LoggedIn {
		return notloginerr()
	}
	return user.Remove()
}

func notloginerr() error {
	return errors.New("you have not login")
}