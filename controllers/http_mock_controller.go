package controllers

import (
	"database/sql"
	"fmt"
	"github.com/jspmarc/mockha/api/service"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HttpMockController struct {
	httpMockService service.HttpMockService
}

func NewMockController(e *echo.Echo, mockService service.HttpMockService, prefix string) *HttpMockController {
	c := new(HttpMockController)
	c.httpMockService = mockService

	group := e.Group(prefix)
	group.POST("/", c.registerMock)
	group.PUT("/:id", c.editMock)
	group.DELETE("/:id", c.deleteMock)
	group.GET("/", c.getMocks)
	group.POST("/:id/execute", c.executeMock)
	group.POST("/execute", c.executeMock)

	return c
}

func (c *HttpMockController) registerMock(ctx echo.Context) error {
	ctx.Logger().Printf("called registerMock with path %s\n", ctx.Path())
	c.httpMockService.RegisterMock(&model.HttpMock{})
	return ctx.String(http.StatusOK, fmt.Sprintf("called registerMock with path %s", ctx.Path()))
}

func (c *HttpMockController) editMock(ctx echo.Context) error {
	ctx.Logger().Printf("called editMock with path %s\n", ctx.Path())
	c.httpMockService.EditMock(&model.HttpMock{})
	return ctx.String(http.StatusOK, fmt.Sprintf("called editMock with path %s", ctx.Path()))
}

func (c *HttpMockController) deleteMock(ctx echo.Context) error {
	ctx.Logger().Printf("called deleteMock with path %s\n", ctx.Path())
	c.httpMockService.DeleteMock(sql.NullString{String: "", Valid: true}, "", constants.HTTP_METHOD_GET)
	return ctx.String(http.StatusOK, fmt.Sprintf("called deleteMock with path %s", ctx.Path()))
}

func (c *HttpMockController) getMocks(ctx echo.Context) error {
	ctx.Logger().Printf("called getMock with path %s\n", ctx.Path())
	c.httpMockService.GetAllMocks()
	return ctx.String(http.StatusOK, fmt.Sprintf("called getMock with path %s", ctx.Path()))
}

func (c *HttpMockController) executeMock(ctx echo.Context) error {
	ctx.Logger().Printf("called getMock with path %s\n", ctx.Path())
	c.httpMockService.ExecuteMock(sql.NullString{String: "", Valid: true}, "", constants.HTTP_METHOD_GET)
	return ctx.String(http.StatusOK, fmt.Sprintf("called getMock with path %s", ctx.Path()))
}

func (c *HttpMockController) Stop() error {
	return nil
}
