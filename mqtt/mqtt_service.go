package mqtt

type MQService interface {
	SendToDevice(deviceId string, command string)
}

type mqService struct {
	client MQClient
}

func NewMQService() MQService {
	service := mqService{}
	service.client = GetInstance()
	return &service
}

func (s *mqService) SendToDevice(deviceId string, command string) {
	s.client.Publish(deviceId, command)
}
