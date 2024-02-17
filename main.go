package main

import (
	"encoding/json"
	"net/http"

	"github.com/MuhammadSyaugi13/golang-todos/database"
	"github.com/labstack/echo"
)

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.POST("/products", func(ctx echo.Context) error {

		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec("INSERT INTO golang_todos (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "ok")

	})

	e.Start(":8080")
}
