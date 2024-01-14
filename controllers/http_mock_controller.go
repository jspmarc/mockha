package controllers

import (
	"database/sql"
	"fmt"
	"github.com/jspmarc/mockha/api/service"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/dto/http_mock"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/utils"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type HttpMockController struct {
	httpMockService service.HttpMockService
}

func NewHttpMockController(e *echo.Echo, mockService service.HttpMockService, prefix string) *HttpMockController {
	c := new(HttpMockController)
	c.httpMockService = mockService

	group := e.Group(prefix)
	group.POST("/", c.registerMock)
	group.PUT("/:id", c.editMock)
	group.DELETE("/:id", c.deleteMock)
	group.GET("/", c.getMocks)

	return c
}

func (c *HttpMockController) Start() error {
	return c.httpMockService.Start()
}

func (c *HttpMockController) registerMock(ctx echo.Context) error {
	log.Info().Msgf("called registerMock with path %s", ctx.Path())

	var req http_mock.CreateRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if req.ResponseCode < 100 || req.ResponseCode >= 600 {
		errMessage := "mock's response code should be between 100 and 599 (inclusive)"
		return utils.ResponseHttp(ctx, http.StatusBadRequest, &errMessage, nil)
	}

	if (req.RequestBody != nil && req.RequestBodyMimeType == nil) ||
		(req.RequestBody == nil && req.RequestBodyMimeType != nil) {

		errMessage := "Invalid request body and MIME type combination"
		return utils.ResponseHttp(ctx, http.StatusBadRequest, &errMessage, nil)
	}

	if req.Method != constants.HTTP_METHOD_GET && req.Method != constants.HTTP_METHOD_HEAD &&
		req.Method != constants.HTTP_METHOD_POST && req.Method != constants.HTTP_METHOD_PUT &&
		req.Method != constants.HTTP_METHOD_DELETE && req.Method != constants.HTTP_METHOD_CONNECT &&
		req.Method != constants.HTTP_METHOD_OPTIONS && req.Method != constants.HTTP_METHOD_TRACE &&
		req.Method != constants.HTTP_METHOD_PATCH {

		errMessage := "Invalid HTTP method"
		return utils.ResponseHttp(ctx, http.StatusBadRequest, &errMessage, nil)
	}

	if mock, err := c.httpMockService.RegisterMock(&req); err != nil {
		errMsg := err.Error()
		return utils.ResponseHttp(ctx, http.StatusInternalServerError, &errMsg, nil)
	} else {
		return utils.ResponseHttp(ctx, http.StatusOK, nil, mock)
	}
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

func (c *HttpMockController) Stop() error {
	return c.httpMockService.Stop()
}
