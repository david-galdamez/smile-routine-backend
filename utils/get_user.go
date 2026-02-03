package utils

import (
	"context"
)

func GetAuthUser(ctx context.Context) (AuthUser, bool) {
	user, ok := ctx.Value("auth_user").(AuthUser)
	return user, ok
}
