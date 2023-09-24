package auth

import (
	"context"

	contextkeyenum "github.com/doyoque/service-a/internal/enum/contextkey"
)

func GetUserAuthorizationContext(ctx context.Context) *UserAuthorization {
	if ctx == nil {
		return nil
	}

	if val, ok := ctx.Value(contextkeyenum.Authorization).(UserAuthorization); ok {
		return &val
	}

	return nil
}
