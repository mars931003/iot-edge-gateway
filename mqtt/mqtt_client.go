package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-edge-gateway/exceptions"
	"iot-edge-gateway/utils"
	"log"
	"time"
)

const mqttServer = "tcp://127.0.0.1:1883"
const gatewayId = "edge-gateway-1"

type MessageBody struct {
	SystemId string      `json:"systemId"`
	Payload  interface{} `json:"payload"`
}

var systemId = ""
var mqttClient mqtt.Client

var Message mqtt.MessageHandler = subscription

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
	options.SetAutoReconnect(true)
	options.SetConnectTimeout(time.Duration(30) * time.Second)
	mqttClient = mqtt.NewClient(options)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Printf("MQTT client was connected : %s", mqttServer)
}

func GetInstance() MQClient {
	return &mqClient{}
}

func (mqClient) getClient() *mqtt.Client {
	return &mqttClient
}

func subscription(client mqtt.Client, message mqtt.Message) {
	body := MessageBody{}
	err := utils.Byte2Model(message.Payload(), &body)
	exceptions.ExceptionHandler(err, "byte convert to model failed !")
	if body.SystemId != systemId {
		return
	}
	log.Println(body)
	//实时接收
	//接收到的消息缓存到redis,响应式消费
}

//func GetMqttClient() mqtt.Client {
//	return s.client
//}

func (s *mqClient) Publish(deviceId string, command string) {
	log.Printf("**** start to send message to topic : %s", deviceId)
	token := mqttClient.Publish(deviceId, 1, false, command)
	if token.Wait() {
		log.Println("waiting for receive ")
	} else {
		log.Println("message send succeed ")
	}
	log.Println("aaa")
}
