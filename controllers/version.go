package controllers

import (
	"github.com/api-pokemon/config"
	"github.com/api-pokemon/helpers"
	"github.com/labstack/echo/v4"
)

func Version(c echo.Context) error {
	res := helpers.Map{
		"version": config.Get("APP_VERSION").String(),
	}

	return helpers.Response(c, 201, res)
}
