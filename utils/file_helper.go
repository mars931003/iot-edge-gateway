package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"iot-edge-gateway/model"
	"log"
)

const FilePath = "device_config.json"

func WriteModelToFile(object interface{}) error {
	bts, err := Model2Byte(object)
	var out bytes.Buffer
	err = json.Indent(&out, bts, "", "\t")
	err = ioutil.WriteFile(FilePath, out.Bytes(), 0666)
	return err
}

func ReadModelFromFile(filename string) (*[]model.DeviceConfig, error) {
	res, err := ioutil.ReadFile(filename)
	var list []model.DeviceConfig
	err = Byte2Model(res, &list)
	log.Println(list)
	return &list, err
}
