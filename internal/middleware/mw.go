package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

const (
	adminUserRole  string = "admin"
	userRoleHeader string = "User-Role"
)

func CheckUserRole(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRoleHeader := ctx.Request().Header[userRoleHeader]
		if isHeaderHasAdminRole(userRoleHeader) {
			log.Print("red button user detected")
		}
		return handler(ctx)
	}
}

func isHeaderHasAdminRole(userRole []string) bool {
	switch {
	case isEmpty(userRole):
		return false
	case userRole[0] == adminUserRole:
		return true
	}
	return false
}

func isEmpty(header []string) bool {
	return len(header) == 0
}
