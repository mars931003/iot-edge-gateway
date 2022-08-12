package mqtt

type MQService interface {
	SendToDevice(deviceId string, command string)
}

type mqService struct {
	client MQClient
}

func NewMQService() MQService {
	return &mqService{client: &mqClient{}}
}

func (s *mqService) SendToDevice(deviceId string, command string) {
	s.client.Publish(deviceId, command)
}
