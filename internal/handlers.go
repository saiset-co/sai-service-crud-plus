package internal

import (
	"encoding/json"
	"github.com/saiset-co/sai-service/middlewares"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/saiset-co/sai-service-crud-plus/internal/actions"
	"github.com/saiset-co/sai-service-crud-plus/types"
	"github.com/saiset-co/sai-service/service"
)

const (
	supportedMethodCreate = "create"
	supportedMethodRead   = "read"
	supportedMethodUpdate = "update"
	supportedMethodDelete = "delete"
)

func (is InternalService) NewHandler() service.Handler {
	return service.Handler{
		supportedMethodCreate: service.HandlerElement{
			Name:        "Create documents",
			Description: "Create documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, supportedMethodCreate)
				if err != nil {
					return nil, http.StatusInternalServerError, err
				}

				return actions.NewSaveAction(is.Storage, is.Prefix).Handle(request)
			},
			Middlewares: []service.Middleware{
				middlewares.CreateAuthMiddleware(is.AuthUrl, is.Name, supportedMethodCreate),
			},
		},
		supportedMethodRead: service.HandlerElement{
			Name:        "Read documents",
			Description: "Read documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, supportedMethodRead)
				if err != nil {
					return nil, http.StatusInternalServerError, err
				}

				return actions.NewGetAction(is.Storage, is.Prefix).Handle(request)
			},
			Middlewares: []service.Middleware{
				middlewares.CreateAuthMiddleware(is.AuthUrl, is.Name, supportedMethodRead),
			},
		},
		supportedMethodUpdate: service.HandlerElement{
			Name:        "Update documents",
			Description: "Update documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, supportedMethodUpdate)
				if err != nil {
					return nil, http.StatusInternalServerError, err
				}

				return actions.NewUpdateAction(is.Storage, is.Prefix).Handle(request)
			},
			Middlewares: []service.Middleware{
				middlewares.CreateAuthMiddleware(is.AuthUrl, is.Name, supportedMethodUpdate),
			},
		},
		supportedMethodDelete: service.HandlerElement{
			Name:        "Delete documents",
			Description: "Delete documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, supportedMethodDelete)
				if err != nil {
					return nil, http.StatusInternalServerError, err
				}

				return actions.NewDeleteAction(is.Storage, is.Prefix).Handle(request)
			},
			Middlewares: []service.Middleware{
				middlewares.CreateAuthMiddleware(is.AuthUrl, is.Name, supportedMethodDelete),
			},
		},
	}
}

func (is InternalService) convertRequest(data interface{}, requestType string) (types.IRequest, error) {
	switch requestType {
	case supportedMethodRead:
		request := types.ReadRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - marshaling - get")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - get")
		}

		err = validator.New().Struct(request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - validation - get")
		}

		return request, nil
	case supportedMethodCreate:
		request := types.CreateRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - marshaling - save")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - save")
		}

		err = validator.New().Struct(request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - validation - save")
		}

		return request, nil
	case supportedMethodUpdate:
		request := types.UpdateRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - marshaling - update")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - update")
		}

		err = validator.New().Struct(request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - validation - update")
		}

		return request, nil
	case supportedMethodDelete:
		request := types.DeleteRequest{}
		dataJson, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - marshaling - delete")
		}

		err = json.Unmarshal(dataJson, &request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - unmarshaling - delete")
		}

		err = validator.New().Struct(request)
		if err != nil {
			return nil, errors.Wrap(err, "convertRequest - validation - delete")
		}

		return request, nil
	}

	return nil, errors.Wrap(errors.New("no variable type"), "convertRequest")
}
