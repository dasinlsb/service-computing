package user

import (
	"agenda/model"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
)

const USER_DB string = "allUsers.txt"
const SESSION_DB string = "curUser.txt"

type User struct {
	Username, Password, Email, Phone string
}

type Session struct {
	LoggedIn bool
	Username string
}

var users []User
var session *Session

var sessionDb *model.Store
var userDb *model.Store

func Load(prefix string) error {
	if err := os.MkdirAll(prefix, os.ModePerm); err != nil {
		return err
	}
	session = &Session{false, ""}
	sessionDb = &model.Store{path.Join(prefix, SESSION_DB)}
	if err := sessionDb.Load(session); err != nil && err != io.EOF {
		return err
	}
	users = []User{}
	userDb = &model.Store{path.Join(prefix, USER_DB)}
	if err := userDb.Load(&users); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func Add(u User) error {
	users = append(users, u)
	return PersistUsers()
}

func Login(username string) error {
	*session =Session{true, username}
	return PersistSession()
}
func Logout() error {
	session.LoggedIn = false
	return PersistSession()
}
func Auth(username, password string) error {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			_ = Login(username)
			return nil
		}
	}
	return errors.New("incorrect username or password, authentication failed")
}

func FindOneByName(username string) (*User, bool) {
	for _, user := range users {
		if user.Username == username {
			return &user, true
		}
	}
	return nil, false
}

func ValidateNewUser(u User) []string {
	sects := []string{"Username", "Email", "Phone"}
	var repeats []string
	for _, sect := range sects {
		find := false
		for _, user := range users {
			if reflect.Indirect(reflect.ValueOf(user)).FieldByName(sect).String() == reflect.Indirect(reflect.ValueOf(u)).FieldByName(sect).String() {
				find = true
				break
			}
		}
		if find {
			repeats = append(repeats, sect)
		}
	}
	return repeats
}

func Remove() error {
	session.LoggedIn = false
	if err := PersistSession(); err != nil {
		return err
	}
	for i, u := range users {
		if u.Username == session.Username {
			users = append(users[:i], users[i+1:]...)
			if err := PersistUsers(); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
	return fmt.Errorf("the username [%v] of current session not found in users database ?!", session.Username)
}

func PersistSession() error {
	return sessionDb.Persist(session)
}
func PersistUsers() error {
	return userDb.Persist(users)
}

func GetSession() (*Session) {
	return session
}

func GetUsers() ([]User) {
	return users
}