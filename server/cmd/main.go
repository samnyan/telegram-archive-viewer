package main

import (
	"github.com/labstack/echo/v4"
	"github.com/samnyan/go-telegram-archive/server/config"
	"github.com/samnyan/go-telegram-archive/server/database"
	"github.com/samnyan/go-telegram-archive/server/logger"
	"github.com/samnyan/go-telegram-archive/server/server"
)

func main() {
	conf := config.LoadConfig("config", "yaml")

	e := echo.New()

	db := database.Initialize(conf)

	server.Initialize(e, db, conf)

	logger.Fatal().Err(e.Start(":" + conf.Server.Port)).Msg("Web Server start failed")
}
