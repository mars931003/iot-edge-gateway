package interf

import (
	"github.com/emicklei/go-restful"
	"io"
	"io/ioutil"
	"iot-edge-gateway/exceptions"
	"iot-edge-gateway/model"
	"iot-edge-gateway/service"
	"iot-edge-gateway/utils"
)

type DeviceApi interface {
	GetDeviceStatusById(request *restful.Request, response *restful.Response)
	GetDevicesStatus(request *restful.Request, response *restful.Response)
	GeneratorConfig(request *restful.Request, response *restful.Response)
	DeviceCommand(request *restful.Request, response *restful.Response)
}

type deviceApi struct {
	deviceService service.DeviceService
}

func NewDeviceApi() DeviceApi {
	controller := deviceApi{}
	controller.deviceService = service.NewDeviceService()
	return &controller
}

func (api *deviceApi) GetDeviceStatusById(request *restful.Request, response *restful.Response) {
	serviceResponse := api.deviceService.GetDeviceStatusById(request.PathParameter("deviceId"))
	writeResponse(response, serviceResponse)
}

func (api *deviceApi) GetDevicesStatus(request *restful.Request, response *restful.Response) {
	writeResponse(response, api.deviceService.GetDeicesStatus())
}

func (api *deviceApi) GeneratorConfig(request *restful.Request, response *restful.Response) {
	body, err := ioutil.ReadAll(request.Request.Body)
	var mod = make([]model.DeviceConfig, 0)
	err = utils.Byte2Model(body, &mod)
	exceptions.ExceptionHandler(err, "unable to get object in request parameter ,get request body error !")
	writeResponse(response, api.deviceService.GenerateConfigFile(mod))
}

func (api *deviceApi) DeviceCommand(request *restful.Request, response *restful.Response) {
	body, err := ioutil.ReadAll(request.Request.Body)
	exceptions.ExceptionHandler(err, "unable to get object in request parameter ,get request body error !")
	writeResponse(response, api.deviceService.SendCommand(request.PathParameter("deviceId"), string(body)))
}

//func SendCommand(deviceId string, command string) {
//	mqtt.SendToDevice(deviceId, command)
//}

func writeResponse(response *restful.Response, apiResponse *model.ApiResponse) {
	str, err := apiResponse.ToString()
	_, err = io.WriteString(response.ResponseWriter, str)
	exceptions.ExceptionHandler(err, "Http response failed, write result exception !")
}

//func readRequestBody(request *restful.Request) ( /*[]model.DeviceConfig*/ interface{}, error) {
//	bts, err := ioutil.ReadAll(request.Request.Body)
//	//requestBody := make([]model.DeviceConfig, 0)
//	result, err := utils.Byte2Any(bts)
//	return result, err
//
//	//err = json.Unmarshal(bts, &requestBody)
//	//return requestBody, err
//}
