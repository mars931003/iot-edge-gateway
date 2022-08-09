package service

import (
	"errors"
	"iot-edge-gateway/model"
	"iot-edge-gateway/mqtt"
	"iot-edge-gateway/utils"
	"log"
)

type DeviceService interface {
	GetDeviceStatusById(id string) *model.ApiResponse
	GetDeicesStatus() *model.ApiResponse
	GenerateConfigFile(config []model.DeviceConfig) *model.ApiResponse
	SendCommand(deviceId string, command string) *model.ApiResponse
}

type deviceService struct {
	mqttService mqtt.MQService
}

func NewDeviceService() DeviceService {
	service := deviceService{}
	service.mqttService = mqtt.NewMQService()
	return &service
}

func (d *deviceService) GetDeviceStatusById(id string) *model.ApiResponse {
	log.Println(id)
	response := model.ApiResponse{}
	//TODO 调用mqtt发送设备指令获取设备状态
	response.Successful("alive")
	return &response
}

func (d *deviceService) GetDeicesStatus() *model.ApiResponse {
	log.Println("get all devices status")
	err := getStatus("获取设备状态")
	response := model.ApiResponse{}
	if err != nil {
		return response.Failed(false).Message(err.Error())
	} else {
		return response.Successful(true)
	}
}

func (d *deviceService) GenerateConfigFile(config []model.DeviceConfig) *model.ApiResponse {
	err := utils.WriteModelToFile(config)
	response := model.ApiResponse{}
	if err != nil {
		return response.Failed(false).Message(err.Error())
	} else {
		return response.Successful(true)
	}
}

func getStatus(commandDescription string) error {
	d := deviceService{}
	list, err := utils.ReadModelFromFile(utils.FilePath)
	var content string
	//获取所有设备的指定命令
	for _, v := range *list {
		content, err = filterCommand(commandDescription, v)
		d.SendCommand(v.DeviceId, content)
	}
	return err
}

func (d *deviceService) SendCommand(deviceId string, command string) *model.ApiResponse {
	log.Println(deviceId)
	log.Println(command)
	d.mqttService.SendToDevice(deviceId, command)
	return &model.ApiResponse{}
}

func filterCommand(commandDesc string, device model.DeviceConfig) (string, error) {
	for _, v := range device.Commands {
		if v.Description == commandDesc {
			return v.Content, nil
		}
	}
	return "", errors.New("no command match")
}
