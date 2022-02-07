package main

import (
	"echoauth/src/controller"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//const jwtSecretKey = "my-super-secret-key"
//
//func main() {
//	e := echo.New()
//
//	// logging and panic recovery middleware
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	// unrestricted route
//	e.GET("/api/unrestricted", unrestricted)
//	// add the login route
//	e.POST("/api/login", login)
//
//	// add a restricted group
//	r := e.Group("/api")
//	// apply the jwt middleware to the route group
//	r.Use(middleware.JWT([]byte(jwtSecretKey)))
//	r.GET("/restricted", restricted)
//
//	e.Logger.Fatal(e.Start(":1323"))
//}
//
//func unrestricted(c echo.Context) error {
//	return c.JSON(http.StatusOK, map[string]string{
//		"message": "Success! The status is 200",
//	})
//}
//func restricted(c echo.Context) error {
//	// do a fancy dance to get the token's email
//	user := c.Get("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	email := claims["email"].(string)
//
//	return c.JSON(http.StatusOK, map[string]string{
//		"message": "hello email address: " + email,
//	})
//}
//
//func login(c echo.Context) error {
//	email := c.FormValue("email")
//	password := c.FormValue("password")
//
//	fmt.Println(email + " : " + password)
//
//	// in our case, the only "valid user and password" is
//	// user: rickety_cricket@example.com pw: shhh!
//	// really, this would be connected to any database and
//	// retrieving the user and validating the password
//	if email != "rickety_cricket@example.com" || password != "shhh!" {
//		return echo.ErrUnauthorized
//	}
//
//	// create token
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// set claims
//	claims := token.Claims.(jwt.MapClaims)
//	// add any key value fields to the token
//	claims["email"] = "rickety_cricket@example.com"
//	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//
//	// generate encoded token and send it as response.
//	t, err := token.SignedString([]byte("secret"))
//	if err != nil {
//		return err
//	}
//
//	// return the token for the consumer to grab and save
//	return c.JSON(http.StatusOK, map[string]string{
//		"token": t,
//	})
//}

const jwtSecretKey = "my-super-secret-key"

func login(c echo.Context) error {
	fmt.Println("email + password")
	email := c.FormValue("email")
	password := c.FormValue("password")

	fmt.Println(email + password)

	// Throws unauthorized error
	if email != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "admin"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	//t, err := token.SignedString([]byte("secret"))
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success! The status is 200",
	})
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello email address: " + email,
	})
}

func main() {
	e := echo.New()
	e.Use(echosession.New())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodDelete},
	}))

	// Login route
	e.POST("/api/login", login)

	// Unauthenticated route
	e.GET("/api/unrestricted", accessible)

	// Restricted group
	r := e.Group("/api/restricted")
	//r.Use(middleware.JWT([]byte("secret")))
	r.Use(middleware.JWT([]byte(jwtSecretKey)))

	r.GET("", restricted)

	namespaceGroup := e.Group("/setting/namespaces")
	//namespaceGroup.Use(middleware.JWT([]byte(jwtSecretKey)))
	//namespaceGroup.Use(middleware.JWT([]byte(os.Getenv("LoginAccessSecret"))))
	namespaceGroup.GET("/namespace/list", controller.GetNameSpaceList)
	namespaceGroup.POST("/namespace/reg/proc", controller.NameSpaceRegProc)
	namespaceGroup.DELETE("/namespace/del/:nameSpaceID", controller.NameSpaceDelProc)

	e.Logger.Fatal(e.Start(":1323"))
}
