package repository

import (
	"database/sql"

	"github.com/MIHAIL33/Ponylab-Go/model"
)

type DeviceStateInterface interface {
	Create(devState model.DeviceState) error 
	GetAll() (*[]model.DeviceState, error)
	GetAllUniq() (*[]model.DeviceState, error)
}

type Repository struct {
	DeviceStateInterface
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DeviceStateInterface: NewDeviceStateRepository(db),
	}
}