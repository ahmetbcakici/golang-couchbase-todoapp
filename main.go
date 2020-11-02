package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

type Cat struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func greetings(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the Echo server!")
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

func main() {
	fmt.Println("Hello World!")

	e := echo.New()

	e.GET("/", greetings)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)

	e.Start(":8000")
}
