package internal

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/saiset-co/sai-service-crud-plus/internal/actions"
	"github.com/saiset-co/sai-service-crud-plus/types"
	"github.com/saiset-co/sai-service/service"
)

func (is InternalService) NewHandler() service.Handler {
	return service.Handler{
		"create": service.HandlerElement{
			Name:        "Create documents",
			Description: "Create documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "save")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewSaveAction(is.Storage, is.Prefix).Handle(request)
			},
		},
		"read": service.HandlerElement{
			Name:        "Read documents",
			Description: "Read documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "get")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewGetAction(is.Storage, is.Prefix).Handle(request)
			},
		},
		"update": service.HandlerElement{
			Name:        "Update documents",
			Description: "Update documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "update")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewUpdateAction(is.Storage, is.Prefix).Handle(request)
			},
		},
		"delete": service.HandlerElement{
			Name:        "Delete documents",
			Description: "Delete documents",
			Function: func(data interface{}, metadata interface{}) (interface{}, int, error) {
				request, err := is.convertRequest(data, "delete")
				if err != nil {
					return nil, 500, err
				}

				return actions.NewDeleteAction(is.Storage, is.Prefix).Handle(request)
			},
		},
	}
}

func (is InternalService) convertRequest(data interface{}, requestType string) (types.IRequest, error) {
	switch requestType {
	case "read":
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
	case "create":
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
	case "update":
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
	case "delete":
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
