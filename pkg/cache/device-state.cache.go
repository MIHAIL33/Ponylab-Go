package cache

import "github.com/MIHAIL33/Ponylab-Go/model"

type CacheDeviceState struct {
	devState map[string]model.DeviceState
}

func NewCacheDeviceState() *CacheDeviceState {
	return &CacheDeviceState{
		devState: make(map[string]model.DeviceState),
	}
}

func (c *CacheDeviceState) AddAll(devS []model.DeviceState) error {
	for _, val := range devS {
		c.devState[val.UID] = val
	}
	return nil
}

func (c *CacheDeviceState) AddOne(devS model.DeviceState) error {
	c.devState[devS.UID] = devS
	return nil
}

func (c *CacheDeviceState) GetAll() (*[]model.DeviceState, error) {
	var res []model.DeviceState
	for _, val := range c.devState {
		res = append(res, val)
	}
	return &res, nil
}

func (c *CacheDeviceState) GetOneByUID(uid string) (*model.DeviceState, error) {
	return nil, nil
}