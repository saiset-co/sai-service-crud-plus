package main

import (
	"github.com/saiset-co/sai-service-crud-plus/internal"
	"github.com/saiset-co/sai-service-crud-plus/logger"
	"github.com/saiset-co/sai-service/service"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

const serviceName = "Crud"

func main() {
	svc := service.NewService(serviceName)

	svc.RegisterConfig("config.yml")

	logger.Logger = svc.Logger

	storageUrl := svc.GetConfig("common.storage.url", "").(string)
	storageToken := svc.GetConfig("common.storage.token", "").(string)

	store := &adapter.SaiStorage{
		Url:   storageUrl,
		Token: storageToken,
	}

	authUrl := svc.GetConfig("common.auth.url", "").(string)

	prefix := svc.GetConfig("common.dictionary_prefix", "").(string)

	is := internal.InternalService{
		Context: svc.Context,
		Name:    serviceName,
		Storage: store,
		Prefix:  prefix,
		AuthUrl: authUrl,
	}

	svc.RegisterHandlers(
		is.NewHandler(),
	)

	svc.Start()
}
