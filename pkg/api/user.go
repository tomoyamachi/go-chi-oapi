package api

import (
	"net/http"
)

type UserService struct{}

func (s UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
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
