package model

import (
	"time"
)

type DeviceState struct {
	UID string `json:"uid"`
	Data string `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}