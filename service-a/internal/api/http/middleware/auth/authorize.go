package auth

import (
	resourceenum "github.com/doyoque/service-a/internal/enum/resource"
)

type UserAuthorization struct {
	Type      string
	UserId    int
	Username  string
	Resources []resourceenum.Resource
}
