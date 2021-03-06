// Package store provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package store

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi"
)

// Address defines model for Address.
type Address struct {
	City   *string `json:"city,omitempty"`
	State  *string `json:"state,omitempty"`
	Street *string `json:"street,omitempty"`
	Zip    *string `json:"zip,omitempty"`
}

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// Category defines model for Category.
type Category struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Customer defines model for Customer.
type Customer struct {
	Address  *[]Address `json:"address,omitempty"`
	Id       *int64     `json:"id,omitempty"`
	Username *string    `json:"username,omitempty"`
}

// Order defines model for Order.
type Order struct {
	Complete *bool      `json:"complete,omitempty"`
	Id       *int64     `json:"id,omitempty"`
	PetId    *int64     `json:"petId,omitempty"`
	Quantity *int32     `json:"quantity,omitempty"`
	ShipDate *time.Time `json:"shipDate,omitempty"`

	// Order Status
	Status *string `json:"status,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	Category  *Category `json:"category,omitempty"`
	Id        *int64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	PhotoUrls []string  `json:"photoUrls"`

	// pet status in the store
	Status *string `json:"status,omitempty"`
	Tags   *[]Tag  `json:"tags,omitempty"`
}

// Tag defines model for Tag.
type Tag struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// User defines model for User.
type User struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  *string `json:"password,omitempty"`
	Phone     *string `json:"phone,omitempty"`

	// User Status
	UserStatus *int32  `json:"userStatus,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// UserArray defines model for UserArray.
type UserArray []User

// PlaceOrderJSONBody defines parameters for PlaceOrder.
type PlaceOrderJSONBody Order

// PlaceOrderRequestBody defines body for PlaceOrder for application/json ContentType.
type PlaceOrderJSONRequestBody PlaceOrderJSONBody

type ServerInterface interface {
	// Returns pet inventories by status (GET /store/inventory)
	GetInventory(w http.ResponseWriter, r *http.Request)
	// Place an order for a pet (POST /store/order)
	PlaceOrder(w http.ResponseWriter, r *http.Request)
	// Delete purchase order by ID (DELETE /store/order/{orderId})
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	// Find purchase order by ID (GET /store/order/{orderId})
	GetOrderById(w http.ResponseWriter, r *http.Request)
}

// GetInventory operation middleware
func GetInventoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, "api_key.Scopes", []string{""})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// PlaceOrder operation middleware
func PlaceOrderCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DeleteOrder operation middleware
func DeleteOrderCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "orderId" -------------
		var orderId int64

		err = runtime.BindStyledParameter("simple", false, "orderId", chi.URLParam(r, "orderId"), &orderId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter orderId: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "orderId", orderId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetOrderById operation middleware
func GetOrderByIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Path parameter "orderId" -------------
		var orderId int64

		err = runtime.BindStyledParameter("simple", false, "orderId", chi.URLParam(r, "orderId"), &orderId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter orderId: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, "orderId", orderId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	r.Group(func(r chi.Router) {
		r.Use(GetInventoryCtx)
		r.Get("/store/inventory", si.GetInventory)
	})
	r.Group(func(r chi.Router) {
		r.Use(PlaceOrderCtx)
		r.Post("/store/order", si.PlaceOrder)
	})
	r.Group(func(r chi.Router) {
		r.Use(DeleteOrderCtx)
		r.Delete("/store/order/{orderId}", si.DeleteOrder)
	})
	r.Group(func(r chi.Router) {
		r.Use(GetOrderByIdCtx)
		r.Get("/store/order/{orderId}", si.GetOrderById)
	})

	return r
}
