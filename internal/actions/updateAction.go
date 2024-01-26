package actions

import (
	"net/http"

	"github.com/saiset-co/sai-service-crud-plus/internal/repositories"
	"github.com/saiset-co/sai-service-crud-plus/types"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

type UpdateAction struct {
	Repository *repositories.Repository
}

func NewUpdateAction(store *adapter.SaiStorage, prefix string) *UpdateAction {
	return &UpdateAction{
		Repository: &repositories.Repository{
			Storage: store,
			Prefix:  prefix,
		},
	}
}

func (action *UpdateAction) Handle(request types.IRequest) (interface{}, int, error) {
	data, errUpd := action.Repository.Update(request.GetCollection(), request.GetSelect(), request.GetData(), request.GetOptions())
	if errUpd != nil {
		return nil, http.StatusInternalServerError, errUpd
	}

	return data, http.StatusOK, nil
}
