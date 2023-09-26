package auth

import (
	resourceenum "github.com/doyoque/service_a/internal/enum/resource"
)

type UserAuthorization struct {
	Type      string
	UserId    int
	Username  string
	Resources []resourceenum.Resource
}
