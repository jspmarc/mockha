package controllers

import (
	"fmt"
	"github.com/josep/mockha/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MockController struct {
	mockService *service.MockService
}

func (c *MockController) registerMock(ctx echo.Context) error {
	ctx.Logger().Printf("called registerMock with path %s\n", ctx.Path())
	c.mockService.RegisterMock()
	return ctx.String(http.StatusOK, fmt.Sprintf("called registerMock with path %s", ctx.Path()))
}

func (c *MockController) editMock(ctx echo.Context) error {
	ctx.Logger().Printf("called editMock with path %s\n", ctx.Path())
	c.mockService.EditMock()
	return ctx.String(http.StatusOK, fmt.Sprintf("called editMock with path %s", ctx.Path()))
}

func (c *MockController) deleteMock(ctx echo.Context) error {
	ctx.Logger().Printf("called deleteMock with path %s\n", ctx.Path())
	c.mockService.DeleteMock()
	return ctx.String(http.StatusOK, fmt.Sprintf("called deleteMock with path %s", ctx.Path()))
}

func (c *MockController) getMock(ctx echo.Context) error {
	ctx.Logger().Printf("called getMock with path %s\n", ctx.Path())
	c.mockService.GetMock()
	return ctx.String(http.StatusOK, fmt.Sprintf("called getMock with path %s", ctx.Path()))
}

func RegisterMockController(e *echo.Echo, prefix string) *MockController {
	c := new(MockController)

	mockServicePort := uint16(8081)
	c.mockService = service.NewMockService(&mockServicePort)
	c.mockService.Start()

	e.POST(fmt.Sprintf("/%s", prefix), c.registerMock)
	e.PATCH(fmt.Sprintf("/%s", prefix), c.editMock)
	e.DELETE(fmt.Sprintf("/%s", prefix), c.deleteMock)
	e.GET(fmt.Sprintf("/%s", prefix), c.getMock)

	return c
}

func (c *MockController) Stop() error {
	return c.mockService.Stop()
}
