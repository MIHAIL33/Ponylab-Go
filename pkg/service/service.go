package service

import (
	"github.com/MIHAIL33/Ponylab-Go/model"
	"github.com/MIHAIL33/Ponylab-Go/pkg/cache"
	"github.com/MIHAIL33/Ponylab-Go/pkg/repository"
)

type DeviceStateInterface interface {
	Create(devState model.DeviceState) error
	GetAll() (*[]model.DeviceState, error)
	GetAllUniq() (*[]model.DeviceState, error)

	AddAllInCache() error
	AddOneInCache(devState model.DeviceState) error
	GetAllFromCache() (*[]model.DeviceState, error)
	GetModelFromCacheById(id string) (*model.DeviceState, error)
}

type Service struct {
	DeviceStateInterface
}

func NewService(repos *repository.Repository, cache * cache.Cache) *Service {
	return &Service{
		DeviceStateInterface: NewDeviceStateService(repos.DeviceStateInterface, cache.DeviceStateInterface),
	}
}