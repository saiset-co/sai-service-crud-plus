package repositories

import (
	"encoding/json"

	"go.uber.org/zap"

	"github.com/saiset-co/sai-service-crud-plus/logger"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

type Repository struct {
	Storage *adapter.SaiStorage
	Prefix  string
}

func (repo *Repository) Create(collection string, documents interface{}) (*adapter.SaiStorageResponse, error) {
	storageRequest := adapter.CreateRequest{
		Collection: repo.Prefix + collection,
		Documents:  documents.([]interface{}),
	}

	return repo.Storage.Send(storageRequest)
}

func (repo *Repository) Read(collection string, selectParams map[string]interface{}, requestOptions interface{}, selectFields []string) (*adapter.SaiStorageResponse, error) {
	options, err := convertToOptions(requestOptions)
	if err != nil {
		logger.Logger.Error("Read", zap.Error(err))
		return nil, err
	}

	storageRequest := adapter.ReadRequest{
		Collection:    repo.Prefix + collection,
		Select:        selectParams,
		Options:       options,
		IncludeFields: selectFields,
	}

	return repo.Storage.Send(storageRequest)
}

func (repo *Repository) Update(collection string, selectParams map[string]interface{}, documentData interface{}, requestOptions interface{}) (*adapter.SaiStorageResponse, error) {
	options, err := convertToOptions(requestOptions)
	if err != nil {
		logger.Logger.Error("Read", zap.Error(err))
		return nil, err
	}

	storageRequest := adapter.UpdateRequest{
		Collection: repo.Prefix + collection,
		Select:     selectParams,
		Document:   documentData,
		Options:    options,
	}

	return repo.Storage.Send(storageRequest)
}

func (repo *Repository) Delete(collection string, selectParams map[string]interface{}) (*adapter.SaiStorageResponse, error) {
	storageRequest := adapter.DeleteRequest{
		Collection: repo.Prefix + collection,
		Select:     selectParams,
	}

	return repo.Storage.Send(storageRequest)
}

func convertToOptions(data interface{}) (*adapter.Options, error) {
	var options = new(adapter.Options)
	optionsBytes, err := json.Marshal(data)
	if err != nil {
		logger.Logger.Error("convertToOptions", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal(optionsBytes, options)
	if err != nil {
		logger.Logger.Error("convertToOptions", zap.Error(err))
		return nil, err
	}

	return options, nil
}
