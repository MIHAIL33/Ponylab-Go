package mqtt

import (
	"errors"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	client mqtt.Client
}

type MQTTPayload struct {
	Topic string
	Payload []byte
}

var messagePubHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	log.Printf("Message %s received on topic %s\n", m.Payload(), m.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
	log.Println("Connected to mqtt")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, err error) {
	log.Printf("Connection lost: %s\n", err.Error())
}

func NewMQTTClient(address, port, clientID string) *MQTTClient {
	broker := "tcp://" + address + ":" + port
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientID)
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler
	client := mqtt.NewClient(options)

	return &MQTTClient{
		client: client,
	}
}

func (m *MQTTClient) Connect() error {
	token := m.client.Connect()
	var wait time.Duration = 1
	var count int = 0

	for {
		if count == 10 {
			return errors.New("no connection to mqtt")
		}
		count++
		if token.Wait() && token.Error() != nil {
			log.Println("No connection to MQTT, repeat...")
			wait *= 2
			
			time.Sleep(wait * time.Second)
			continue
		}
		return nil
	}
}

func (m *MQTTClient) Listen(topic string, ch chan MQTTPayload) {
	token := m.client.Subscribe(topic, 1, func(c mqtt.Client, m mqtt.Message) {
		var mqttPayload MQTTPayload
		mqttPayload.Topic = m.Topic()
		mqttPayload.Payload = m.Payload()
		ch <- mqttPayload
	}) 
	token.Wait()
	log.Println("Subscribed to topic ", topic)
}

func (m *MQTTClient) Disconnect(time uint) {
	m.client.Disconnect(time)
}
