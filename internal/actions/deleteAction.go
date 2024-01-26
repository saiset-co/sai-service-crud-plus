package actions

import (
	"net/http"

	"github.com/saiset-co/sai-service-crud-plus/internal/repositories"
	"github.com/saiset-co/sai-service-crud-plus/types"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

type DeleteAction struct {
	Repository *repositories.Repository
}

func NewDeleteAction(store *adapter.SaiStorage, prefix string) *DeleteAction {
	return &DeleteAction{
		Repository: &repositories.Repository{
			Storage: store,
			Prefix:  prefix,
		},
	}
}

func (action *DeleteAction) Handle(request types.IRequest) (interface{}, int, error) {
	_, errDelete := action.Repository.Delete(request.GetCollection(), request.GetSelect())
	if errDelete != nil {
		return nil, http.StatusInternalServerError, errDelete
	}

	return "Document has been deleted", http.StatusOK, nil
}
