package api

import (
	"net/http"
)

type StoreService struct{}

func (s StoreService) GetInventory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get inventory"))
}
func (s StoreService) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("place order"))
}
func (s StoreService) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete order"))
}
func (s StoreService) GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get order"))
}
