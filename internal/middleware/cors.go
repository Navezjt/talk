package middleware

import "github.com/labstack/echo/v4/middleware"

var AllowAllCors = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins:  []string{"*"},
	AllowHeaders:  []string{"*", "Authorization"},
	AllowMethods:  []string{"*"},
	ExposeHeaders: []string{"*"},
})
