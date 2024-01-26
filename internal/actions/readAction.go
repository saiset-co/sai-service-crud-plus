package actions

import (
	"net/http"

	"github.com/saiset-co/sai-service-crud-plus/internal/repositories"
	"github.com/saiset-co/sai-service-crud-plus/types"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

type ReadAction struct {
	Repository *repositories.Repository
}

func NewGetAction(store *adapter.SaiStorage, prefix string) *ReadAction {
	return &ReadAction{
		Repository: &repositories.Repository{
			Storage: store,
			Prefix:  prefix,
		},
	}
}

func (action *ReadAction) Handle(request types.IRequest) (interface{}, int, error) {
	responseData, errGet := action.Repository.Read(request.GetCollection(), request.GetSelect(), request.GetOptions(), request.GetIncludeFields())
	if errGet != nil {
		return nil, http.StatusInternalServerError, errGet
	}

	return responseData, http.StatusOK, nil
}
