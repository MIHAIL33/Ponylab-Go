package receiver

import (
	"log"
	"strings"
	"time"

	"github.com/MIHAIL33/Ponylab-Go/model"
	"github.com/MIHAIL33/Ponylab-Go/pkg/mqtt"
	"github.com/MIHAIL33/Ponylab-Go/pkg/service"
)

type Receiver struct {
	service *service.Service
	ch chan mqtt.MQTTPayload
}

func NewReceiver(service *service.Service, ch chan mqtt.MQTTPayload) *Receiver {
	return &Receiver{
		service: service,
		ch: ch,
	}
}

func (r *Receiver) GetMQTTPayload() {
	var mqttPayload mqtt.MQTTPayload
	var devS model.DeviceState
	for {
		mqttPayload = <-r.ch
		
		devS = convert(mqttPayload)

		err := r.service.Create(devS)
		if err != nil {
			log.Println(err.Error())
		}

		err = r.service.AddOneInCache(devS)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func convert(mqttPayload mqtt.MQTTPayload) model.DeviceState {
	var devS model.DeviceState

	devS.UID = strings.Split(mqttPayload.Topic, "/")[2] //serial number
	devS.Data = string(mqttPayload.Payload)
	devS.CreatedAt = time.Now()

	return devS
}