package main

import (
	"iot-edge-gateway/commons"
	service "iot-edge-gateway/interf"
	_ "iot-edge-gateway/mqtt"
)

func main() {
	registry := commons.Registrar{}
	api := service.NewDeviceApi()
	registry.RegisterRoute("/device/status/load/id/{deviceId}", commons.Get, api.GetDeviceStatusById)
	registry.RegisterRoute("/device/status/load_all", commons.Get, api.GetDevicesStatus)
	registry.RegisterRoute("/device_config/generate/config", commons.Post, api.GeneratorConfig)
	registry.RegisterRoute("/device_command/device_id/{deviceId}", commons.Post, api.DeviceCommand)
	commons.ApplicationRun(8090)
}
