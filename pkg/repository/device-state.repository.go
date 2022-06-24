package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MIHAIL33/Ponylab-Go/model"
)

type DeviceStateRepository struct {
	db *sql.DB
}

func NewDeviceStateRepository(db *sql.DB) *DeviceStateRepository {
	return &DeviceStateRepository{
		db: db,
	}
}

func (d *DeviceStateRepository) Create(devState model.DeviceState) error {
	createQuery := fmt.Sprintf("INSERT INTO %s (uid, data, created_at) VALUES ($1, $2, $3)" ,devicesStateTable)
	_, err := d.db.Exec(createQuery, devState.UID, devState.Data, devState.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceStateRepository) GetAll() (*[]model.DeviceState, error) {
	return nil, nil
} 

func (d *DeviceStateRepository) GetAllUniq() (*[]model.DeviceState, error) {
	getQuery := fmt.Sprintf("SELECT DISTINCT on (uid) uid, data, created_at FROM %s ORDER BY uid, created_at desc;", devicesStateTable)
	var res []model.DeviceState
	rows, err := d.db.QueryContext(context.Background(), getQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp model.DeviceState
		err := rows.Scan(
			&temp.UID,
			&temp.Data,
			&temp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, temp)
	}

	return &res, nil
} 
