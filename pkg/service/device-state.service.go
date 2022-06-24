package service

import (
	"github.com/MIHAIL33/Ponylab-Go/model"
	"github.com/MIHAIL33/Ponylab-Go/pkg/cache"
	"github.com/MIHAIL33/Ponylab-Go/pkg/repository"
)

type DeviceStateService struct {
	repos repository.DeviceStateInterface
	cache cache.DeviceStateInterface
}

func NewDeviceStateService(repos repository.DeviceStateInterface, cache cache.DeviceStateInterface) *DeviceStateService {
	return &DeviceStateService{
		repos: repos,
		cache: cache,
	}
}

func (d *DeviceStateService) Create(devState model.DeviceState) error {
	err := d.repos.Create(devState)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceStateService) GetAll() (*[]model.DeviceState, error) {
	return nil, nil
}

func (d *DeviceStateService) GetAllUniq() (*[]model.DeviceState, error) {
	devS, err := d.repos.GetAllUniq()
	if err != nil {
		return nil, err
	}
	return devS, nil
}

func (d *DeviceStateService) AddAllInCache() error {
	devS, err := d.GetAllUniq()
	if err != nil {
		return err
	}
	err = d.cache.AddAll(*devS)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceStateService) AddOneInCache(devState model.DeviceState) error {
	err := d.cache.AddOne(devState)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceStateService) GetAllFromCache() (*[]model.DeviceState, error) {
	devS, err := d.cache.GetAll()
	if err != nil {
		return nil, err
	}
	return devS, nil
}

func (d *DeviceStateService) GetModelFromCacheById(id string) (*model.DeviceState, error) {
	return nil, nil
}