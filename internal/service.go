package internal

import (
	"github.com/saiset-co/sai-service/service"
	"github.com/saiset-co/sai-storage-mongo/external/adapter"
)

type InternalService struct {
	Name    string
	Prefix  string
	Context *service.Context
	Storage *adapter.SaiStorage
	AuthUrl string
}
