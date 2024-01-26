package actions

import (
	"net/http"

	"github.com/saiset-co/sai-service-crud-plus/internal/repositories"
	"github.com/saiset-co/sai-service-crud-plus/types"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

type CreateAction struct {
	Repository *repositories.Repository
}

func NewSaveAction(store *adapter.SaiStorage, prefix string) *CreateAction {
	return &CreateAction{
		Repository: &repositories.Repository{
			Storage: store,
			Prefix:  prefix,
		},
	}
}

func (action *CreateAction) Handle(request types.IRequest) (interface{}, int, error) {
	data, errCreation := action.Repository.Create(request.GetCollection(), request.GetData())
	if errCreation != nil {
		return nil, http.StatusInternalServerError, errCreation
	}

	return data, http.StatusOK, nil
}
