package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-edge-gateway/utils"
	"log"
	"time"
)

const mqttServer = "tcp://127.0.0.1:1883"
const gatewayId = "edge-gateway-1"

type MessageType int

const (
	PublishResponse MessageType = iota
	DeviceSender
)

type MessageBody struct {
	SystemId    string      `json:"systemId"`
	Data        interface{} `json:"data"`
	MessageType MessageType `json:"type"`
}

var (
	mqttClient mqtt.Client
	systemId   = ""
	ch         = make(chan string)

	messageHandler mqtt.MessageHandler = subscription

	connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		log.Printf("mqtt client connected , host : %s", mqttServer)
	}

	connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		log.Printf("mqtt client was lost connected ! error : %s", err.Error())
	}
)

type MQClient interface {
	Publish(deviceId string, command string)
}

type mqClient struct {
}

func init() {
	options := mqtt.NewClientOptions().AddBroker(mqttServer)
	options.SetClientID(gatewayId)
	options.SetCleanSession(false)
	options.SetKeepAlive(1000)
	options.SetDefaultPublishHandler(messageHandler)
	options.SetOnConnectHandler(connectHandler)
	options.SetConnectionLostHandler(connectLostHandler)
	options.SetAutoReconnect(true)
	options.SetConnectTimeout(time.Duration(30) * time.Second)
	mqttClient = mqtt.NewClient(options)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	mqttClient.Subscribe("java-sender", 1, messageHandler)
}

func (s *mqService) GetInstance() MQClient {
	return &mqClient{}
}

func subscription(client mqtt.Client, message mqtt.Message) {
	body := MessageBody{}
	err := utils.Byte2Model(message.Payload(), &body)
	if err != nil {
		log.Printf("malformed message, can not handle message, message sent exception ! %s", err.Error())
		return
	}
	switch body.MessageType {
	case PublishResponse:
		ch <- string(message.Payload())
		log.Printf("receive message from %s , message is : %s", message.Topic(), string(message.Payload()))
		return
	case DeviceSender:
		return
	}
	//接收到的消息通过channel发送到主线程进行响应,响应式消费
}

func (s *mqClient) Publish(deviceId string, command string) {
	token := mqttClient.Publish(deviceId, 1, false, command)
	if _, ok := <-ch; token.Wait() && token.Error() != nil && !ok {
		log.Printf("send message 【%s】 to topic 【%s】 failed , ", command, deviceId)
	} else {
		log.Printf("send message 【%s】 to topic 【%s】 succeed ", command, deviceId)
	}
}
