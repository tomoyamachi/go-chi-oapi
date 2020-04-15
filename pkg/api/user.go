package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tomoyamachi/chi-oapi/pkg/gen/user"
)

type UserService struct{}

func responseCommon(w http.ResponseWriter, result interface{}, errP *error) {
	err := *errP
	log.Print(result, errP)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	if result != nil {
		json.NewEncoder(w).Encode(result)
		return
	}
}

func (s UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	var result interface{}
	var err error
	defer responseCommon(w, &result, &err)
	var u user.CreateUserJSONBody
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		return
	}

	id := int64(1)
	result = user.User{
		Email: u.Email,
		Id:    &id,
	}
}
func (s UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
func (s UserService) CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create list with input"))
}
func (s UserService) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login user"))
}
func (s UserService) LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout user"))
}
func (s UserService) GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user by name"))
}
func (s UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}
