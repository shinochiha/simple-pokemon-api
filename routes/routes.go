package routes

import (
	"github.com/api-pokemon/controllers"
	"github.com/api-pokemon/middlewares"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Routes(db *gorm.DB) *echo.Echo {

	e := echo.New()
	e.HTTPErrorHandler = middlewares.ErrorHandler
	e.Use(middlewares.TransactionHandler(db))
	// Auth
	e.POST("/api/v1/register", controllers.Register)
	e.POST("/api/v1/login", controllers.Login)
	e.POST("/api/v1/forgot_password", controllers.ForgotPassword)
	e.POST("/api/v1/reset_password", controllers.ResetPassword)
	e.GET("/api/v1/reset_password/:id", controllers.ResetPassword)
	e.GET("/api/v1/emailver/:id", controllers.EmailVer)
	e.GET("/api/v1/version", controllers.Version)

	e.GET("/api/v1/pokemons", controllers.Pokemon)
	e.GET("/api/v1/pokemons/:id", controllers.PokemonGetById)

	g := e.Group("/api/v1")
	g.Use(middlewares.JWTMiddleware)
	e.GET("/api/v1/logout", controllers.Logout)

	g.GET("/mypokemons", controllers.MyPokemon)
	g.GET("/mypokemons/:id", controllers.MyPokemonGetById)
	g.POST("/capture_pokemon/:id", controllers.HandleCatch)
	g.PUT("/save_pokemon/:id", controllers.SavePokemon)
	g.PUT("/change_name_pokemon/:id", controllers.ChangeNamePokemon)
	g.DELETE("/release_pokemon/:id", controllers.RealeasePokemon)

	return e
}
