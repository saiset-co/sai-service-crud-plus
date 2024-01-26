package main

import (
	"github.com/saiset-co/sai-service-crud-plus/internal"
	"github.com/saiset-co/sai-service-crud-plus/logger"
	"github.com/saiset-co/sai-service/service"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

func main() {
	name := "SaiServiceCrudPlus"

	svc := service.NewService(name)

	svc.RegisterConfig("config.yml")

	logger.Logger = svc.Logger

	storageUrl := svc.GetConfig("common.storage.url", "").(string)
	storageToken := svc.GetConfig("common.storage.token", "").(string)

	store := &adapter.SaiStorage{
		Url:   storageUrl,
		Token: storageToken,
	}

	prefix := svc.GetConfig("common.dictionary_prefix", "").(string)

	is := internal.InternalService{
		Context: svc.Context,
		Name:    name,
		Storage: store,
		Prefix:  prefix,
	}

	svc.RegisterHandlers(
		is.NewHandler(),
	)

	svc.Start()
}
