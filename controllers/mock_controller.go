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

type MockController struct {
	mockService service.MockService
}

func NewMockController(e *echo.Echo, mockService service.MockService, prefix string) *MockController {
	c := new(MockController)
	c.mockService = mockService

	group := e.Group(prefix)
	group.POST("/", c.registerMock)
	group.PUT("/:id", c.editMock)
	group.DELETE("/:id", c.deleteMock)
	group.GET("/", c.getMocks)
	group.POST("/:id/execute", c.executeMock)
	group.POST("/execute", c.executeMock)

	return c
}

func (c *MockController) registerMock(ctx echo.Context) error {
	ctx.Logger().Printf("called registerMock with path %s\n", ctx.Path())
	c.mockService.RegisterMock(&model.HttpMock{})
	return ctx.String(http.StatusOK, fmt.Sprintf("called registerMock with path %s", ctx.Path()))
}

func (c *MockController) editMock(ctx echo.Context) error {
	ctx.Logger().Printf("called editMock with path %s\n", ctx.Path())
	c.mockService.EditMock(&model.HttpMock{})
	return ctx.String(http.StatusOK, fmt.Sprintf("called editMock with path %s", ctx.Path()))
}

func (c *MockController) deleteMock(ctx echo.Context) error {
	ctx.Logger().Printf("called deleteMock with path %s\n", ctx.Path())
	c.mockService.DeleteMock(sql.NullString{String: "", Valid: true}, "", constants.HTTP_METHOD_GET)
	return ctx.String(http.StatusOK, fmt.Sprintf("called deleteMock with path %s", ctx.Path()))
}

func (c *MockController) getMocks(ctx echo.Context) error {
	ctx.Logger().Printf("called getMock with path %s\n", ctx.Path())
	c.mockService.GetAllMocks()
	return ctx.String(http.StatusOK, fmt.Sprintf("called getMock with path %s", ctx.Path()))
}

func (c *MockController) executeMock(ctx echo.Context) error {
	ctx.Logger().Printf("called getMock with path %s\n", ctx.Path())
	c.mockService.ExecuteMock(sql.NullString{String: "", Valid: true}, "", constants.HTTP_METHOD_GET)
	return ctx.String(http.StatusOK, fmt.Sprintf("called getMock with path %s", ctx.Path()))
}

func (c *MockController) Stop() error {
	return nil
}
