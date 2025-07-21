package server

import (
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samnyan/go-telegram-archive/server/config"
	"github.com/samnyan/go-telegram-archive/server/database"
)

func Initialize(e *echo.Echo, db *database.DB, conf *config.Config) {
	// Middleware
	e.Use(middleware.Recover())
	e.Use(RequestLogger())

	e.Use(middleware.CORS())
	e.HTTPErrorHandler = ErrorHandler

	// Jwt
	secret := conf.Jwt.Secret
	jwtConfig := echojwt.Config{
		SigningKey:  []byte(secret),
		TokenLookup: "header:Authorization:Bearer ",
	}

	// Router
	g := e.Group("/api/go")
	g.Use(echojwt.WithConfig(jwtConfig))
}
