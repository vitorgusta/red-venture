package main

import (
	"net/http"

	"github.com/vitorgusta/red-venture/api/controller"

	auth "github.com/vitorgusta/red-venture/api/authentication"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// //CORS
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	// }))

	// Login route
	e.POST("/login", auth.Login)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	//users
	r.GET("/users", controller.FindAll)
	r.GET("/user/:id", controller.FindById)
	r.POST("/user", controller.CreateUser)

	//widgets
	r.GET("/widgets", controller.FindAllWidget)
	r.GET("/widgets/:id", controller.FindWidgetById)
	r.POST("/widgets", controller.CreateWidget)
	r.PUT("/widgets/:id", controller.UpdateWidget)
	e.Logger.Fatal(e.Start(":1323"))
}
