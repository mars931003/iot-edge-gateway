package main

import (
	"iot-edge-gateway/commons"
	service "iot-edge-gateway/interf"
	_ "iot-edge-gateway/mqtt"
)

const systemId = true

func main() {
	registry := commons.Registrar{}
	api := service.NewDeviceApi()
	registry.RegisterRoute("/device/status/load/id/{deviceId}", commons.Get, api.GetDeviceStatusById)
	registry.RegisterRoute("/device/status/load_all", commons.Get, api.GetDevicesStatus)
	registry.RegisterRoute("/device_config/generate/config", commons.Post, api.GeneratorConfig)
	registry.RegisterRoute("/device_command/device_id/{deviceId}", commons.Post, api.DeviceCommand)
	commons.ApplicationRun(8090)
}

//type ResBody struct {
//	Aaa string `json:"aaa"`
//}
//
//func test(request *restful.Request, response *restful.Response) {
//	res, err := ioutil.ReadAll(request.Request.Body)
//	b := ResBody{}
//	err = json.Unmarshal(res, &b)
//	if err != nil {
//		return
//	}
//	log.Println(b.Aaa)
//	m := make(map[string]string)
//	m["ccc"] = "abcderf"
//	marshal, err := json.Marshal(m)
//	if err != nil {
//		return
//	}
//	_, err = io.WriteString(response.ResponseWriter, string(marshal))
//	if err != nil {
//		log.Fatal("")
//	}
//}
