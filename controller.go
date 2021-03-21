package auth

import (
	"encoding/json"

	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

type controller struct {
	config *AuthConfig
	model  *model
}

func newController(config *AuthConfig, model *model) *controller {
	return &controller{
		config: config,
		model:  model,
	}
}

func (c *controller) getSession(ctx *web.Context) error {
	request := &getSessionRequest{}

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

	response, err := c.model.getSession(request)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *controller) refreshSession(ctx *web.Context) error {
	request := &refreshSessionRequest{
		Authorization: ctx.Request.GetHeader(web.HeaderAuthorization),
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	response, err := c.model.refreshToken(request)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *controller) signUp(ctx *web.Context) error {
	request := &signUpRequest{}

	err := ctx.Request.Bind(&request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	response, err := c.model.signUp(request)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.JSON(web.StatusCreated, response)
}

func (c *controller) deactivateUser(ctx *web.Context) error {
	request := &updateUserStatusRequest{
		IdUser: ctx.Request.GetUrlParam("id_user"),
	}

	err := json.Unmarshal(ctx.Request.Body, request)
	if err != nil {
		return ctx.Response.JSON(web.StatusBadRequest, err)
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		return ctx.Response.JSON(web.StatusBadRequest, errs)
	}

	err = c.model.updateUserStatus(request.IdUser, false)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError, ErrorResponse{Code: web.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.Response.NoContent(web.StatusNoContent)
}
