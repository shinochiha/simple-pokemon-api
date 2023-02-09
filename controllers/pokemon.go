package controllers

import (
	"github.com/api-pokemon/helpers"
	"github.com/api-pokemon/models"
	"github.com/labstack/echo/v4"
)

func Pokemon(c echo.Context) error {
	m := models.Pokemon{}
	res := m.GetPaginated(helpers.SetContext(c), c.QueryParams())
	return helpers.Response(c, 200, res)
}

func PokemonGetById(c echo.Context) error {
	m := models.PokemonDetail{}
	res := m.GetById(helpers.SetContext(c), c.Param("id"), c.QueryParams())
	return helpers.Response(c, 200, res)
}

func MyPokemon(c echo.Context) error {
	m := models.MyPokemon{}
	res := m.GetPaginated(helpers.SetContext(c), c.QueryParams())
	return helpers.Response(c, 200, res)
}

func MyPokemonGetById(c echo.Context) error {
	m := models.MyPokemon{}
	res := m.GetById(helpers.SetContext(c), c.Param("id"), c.QueryParams())
	return helpers.Response(c, 200, res)
}

func HandleCatch(c echo.Context) error {
	o := new(models.MyPokemon)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	res := o.Create(helpers.SetContext(c), c.Param("id"))
	return helpers.Response(c, 201, res)
}

func SavePokemon(c echo.Context) error {
	o := new(models.MyPokemon)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	o.ID = helpers.Convert(c.Param("id")).Int()
	res := o.SavePokemon(helpers.SetContext(c))
	return helpers.Response(c, 200, res)
}

func ChangeNamePokemon(c echo.Context) error {
	o := new(models.MyPokemon)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	o.ID = helpers.Convert(c.Param("id")).Int()
	res := o.ChangeNamePokemon(helpers.SetContext(c))
	return helpers.Response(c, 200, res)
}

func RealeasePokemon(c echo.Context) error {
	o := new(models.MyPokemon)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	o.ID = helpers.Convert(c.Param("id")).Int()
	res := o.RealeasePokemon(helpers.SetContext(c))
	return helpers.Response(c, 200, res)
}
