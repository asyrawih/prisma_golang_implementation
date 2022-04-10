package router

import (
	todo "github.com/hananloser/prismago/handler/Todo"
	"github.com/hananloser/prismago/prisma/db"
	"github.com/labstack/echo/v4"
)

func TodoRouter(e *echo.Echo, dbClient *db.PrismaClient) {
	todoHandler := todo.NewTodoHandler(dbClient)

	g := e.Group("todos")

	g.POST("", todoHandler.Add)
	g.GET("", todoHandler.ShowAll)
	g.GET("/:id", todoHandler.Find) // Hardcode On Handler
}
