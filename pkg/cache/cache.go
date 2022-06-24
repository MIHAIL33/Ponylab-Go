package cache

import (
	"github.com/MIHAIL33/Ponylab-Go/model"
)

type DeviceStateInterface interface {
	AddAll(devState []model.DeviceState) error
	AddOne(devState model.DeviceState) error
	GetAll() (*[]model.DeviceState, error)
	GetOneByUID(uid string) (*model.DeviceState, error)
}

type Cache struct {
	DeviceStateInterface
}

func NewCache() *Cache {
	return &Cache{
		DeviceStateInterface: NewCacheDeviceState(),
	}
}