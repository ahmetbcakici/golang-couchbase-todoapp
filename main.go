package main

import (
	"encoding/json"
	"fmt"
	"github.com/couchbase/gocb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Cat struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type User struct {
	Id string `json:"uid"`
	Email string `json:"email"`
	Interests []string `json:"interests"`
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("catName")
	catAge := c.QueryParam("catAge")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your cat name is: %s and it's age is: %s", catName, catAge))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"catName": catName,
			"catAge":  catAge,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "You can just send your request with type string or json",
	})
}

func addCat(c echo.Context) error {
	cat := Cat{}

	// defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed on reading the request body! %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(body, &cat)
	if err != nil {
		log.Printf("Failed on unmarshaling! %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your cat %+v", cat)
	return c.String(http.StatusOK, "We got your cat!")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check username and password inside database instead of that
	if username == "admin" && password == "123" {
		cookie := &http.Cookie{}

		// this is the same
		//cookie := new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		return c.String(http.StatusOK, "You were logged in!")
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "horay you are on the secret amdin main page!")
}

func main() {
	fmt.Println("Hello World!")

	e := echo.New()
	e.Use(setServerHeader)

	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.Logger())

	/*
		adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			// check in the DB
			if username == "ahmet" && password == "123" {
				return true, nil
			}

			return false, nil
		}))

	*/

	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "USERNAME",
		Password: "PASSWORD",
	})
	bucket, _ := cluster.OpenBucket("bucketname", "")

	bucket.Manager("", "").CreatePrimaryIndex("", true, false)

	bucket.Upsert("u:kingarthur",
		User{
			Id: "kingarthur",
			Email: "kingarthur@couchbase.com",
			Interests: []string{"Holy Grail", "African Swallows"},
		}, 0)

	// Get the value back
	var inUser User
	bucket.Get("u:kingarthur", &inUser)
	fmt.Printf("User: %v\n", inUser)

	// Use query
	query := gocb.NewN1qlQuery("SELECT * FROM bucketname WHERE $1 IN interests")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"African Swallows"})
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %v", row)
	}

	adminGroup.GET("/main", mainAdmin, checkCookie)

	e.GET("/login", login)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)



	e.Start(":8000")
}

// MIDDLEWARE

func setServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "ABC-Server/1.0")

		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have any cookie")
			}

			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
	}
}
