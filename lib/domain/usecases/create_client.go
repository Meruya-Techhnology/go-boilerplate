package usecases

import (
	"errors"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	htt "github.com/Meruya-Technology/go-boilerplate/lib/common/https"
	rep "github.com/Meruya-Technology/go-boilerplate/lib/domain/repositories"
	req "github.com/Meruya-Technology/go-boilerplate/lib/presentation/schemes/requests"
	ech "github.com/labstack/echo/v4"
)

type CreateClient struct {
	Repository rep.ClientRepository
	Config     cfg.Config
}

// CreateClient example
// @Description An API for create new client
// @ID create-client
// @tags Client
// @Accept  json
// @Produce  json
// @Param payload body requests.CreateClientRequest true "Client payload"
// @Success 200 {object} base.SuccessResponse{data=entities.Client} "Success response"
// @Success 500 {object} base.InternalServerError "Internal Server Error"
// @Success 400 {object} base.BadRequestError "Bad Request"
// @Success 401 {object} base.UnauthorizedError "Unauthorized"
// @Success 403 {object} base.ForbidenError "Forbiden"
// @Success 404 {object} base.NotFoundError "Not Found"
// @Router /client/create [post]
func (c CreateClient) Execute(ctx ech.Context) error {
	/// Compile request
	request := req.CreateClientRequest{}
	err := htt.ParsingParameter(ctx, &request)
	if err != nil {
		return htt.ErrorParsing(ctx, err, nil)
	}

	/// Request validation
	err = c.validate(ctx, request)
	if err != nil {
		return htt.ErrorBadRequest(ctx, err, nil)
	}

	/// Build & run usecase
	result, err := c.build(ctx, request)
	if err != nil {
		return htt.ErrorInternalServerResponse(ctx, err, nil)
	}

	/// Return final result
	return htt.CreatedResponse(ctx, "Client created successfuly", result)
}

func (c CreateClient) validate(ctx ech.Context, Request req.CreateClientRequest) error {
	config := c.Config
	SecretKey := config.Secret
	if Request.Secret != SecretKey {
		return errors.New("Invalid secret key")
	}
	return nil
}

func (c CreateClient) build(ctx ech.Context, Request req.CreateClientRequest) (interface{}, error) {
	repository := c.Repository
	result, err := repository.Create(ctx, Request)
	return result, err
}
