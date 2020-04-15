package api

import (
	"github.com/tomoyamachi/chi-oapi/pkg/gen/store"
	"github.com/tomoyamachi/chi-oapi/pkg/gen/user"
)

var _ store.ServerInterface = StoreService{}
var _ user.ServerInterface = UserService{}
