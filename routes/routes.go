package routes

import (
	"day-13-orm/configs"
	c "day-13-orm/controllers"
	"day-13-orm/middlewares"
	r "day-13-orm/repositories"
	s "day-13-orm/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	DB = configs.InitDB()
	userR = r.NewUserRepository(DB)
	userS = s.NewUserService(userR)
	userC = c.NewUserController(userS)
)

func New() *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()
	
	middlewares.LoggerMiddleware(e)
	// user routing
	e.POST("/users", userC.CreateController)
	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	e.GET("/users", userC.GetUsersController)
	e.GET("/users/:id", userC.GetUserController)
	// auth.DELETE("/users/:id", c.DeleteUserController)
	// auth.PUT("/users/:id", c.UpdateUserController)

	// e.GET("/books", c.GetBooksController)
	// e.GET("/books/:id", c.GetBookController)
	// auth.POST("/books", c.CreateBookController)
	// auth.DELETE("/books/:id", c.DeleteBookController)
	// auth.PUT("/books/:id", c.UpdateBookController)

	return e
}
