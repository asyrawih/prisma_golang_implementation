package todo

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hananloser/prismago/prisma/db"
	"github.com/labstack/echo/v4"
)

type TodoHandlerContract interface {
	Add(c echo.Context) error
	ShowAll(c echo.Context) error
	Find(c echo.Context) error
}

type TodoHandler struct {
	client *db.PrismaClient
}

type TodoResponse struct {
	todo db.TodosModel
}

func NewTodoHandler(client *db.PrismaClient) *TodoHandler {
	return &TodoHandler{client: client}
}

func (todohandler *TodoHandler) Add(c echo.Context) error {
	var todo db.TodosModel // Get Model From Todo

	if err := c.Bind(&todo); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	todos, err := todohandler.client.Todos.CreateOne(db.Todos.Name.Set(todo.Name)).Exec(context.Background())

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func (todohandler *TodoHandler) ShowAll(c echo.Context) error {

	tm, err := todohandler.client.Todos.FindMany().Exec(context.Background())

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tm)

}

func (todoHandler *TodoHandler) Find(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	tm, err := todoHandler.client.Todos.FindUnique(db.Todos.ID.Equals(id)).Exec(context.Background())
	if err != nil {
		c.Logger().Debug(err.Error())
	}
	return c.JSON(http.StatusOK, tm)
}
