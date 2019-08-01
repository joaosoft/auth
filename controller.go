package auth

import (
	"encoding/json"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type Controller struct {
	config     *AuthConfig
	interactor *Interactor
}

func NewController(config *AuthConfig, interactor *Interactor) *Controller {
	return &Controller{
		config:     config,
		interactor: interactor,
	}
}

func (c *Controller) GetSessionHandler(ctx *web.Context) error {
	request := &GetSessionRequest{}

	err := ctx.Request.BindParams(&request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	err = ctx.Request.Bind(&request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	response, err := c.interactor.GetSession(request)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *Controller) RefreshSessionHandler(ctx *web.Context) error {
	request := &RefreshSessionRequest{
		Authorization: ctx.Request.GetHeader(web.HeaderAuthorization),
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	response, err := c.interactor.RefreshToken(request)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *Controller) SignUpHandler(ctx *web.Context) error {
	request := &SignUpRequest{}

	err := ctx.Request.Bind(&request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	response, err := c.interactor.SignUp(request)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusCreated, response)
}

func (c *Controller) DeactivateUserHandler(ctx *web.Context) error {
	request := &ChangeUserStatusRequest{
		IdUser: ctx.Request.GetUrlParam("id_user"),
	}

	err := json.Unmarshal(ctx.Request.Body, request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	err = c.interactor.ChangeUserStatus(request.IdUser, false)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.NoContent(web.StatusNoContent)
}
