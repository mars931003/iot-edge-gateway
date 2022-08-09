package utils

import (
	"encoding/json"
)

type anything = interface{}

//func Byte2Model(data []byte, v anything) (err error, errMsg string) {
//	err = json.Unmarshal(data, v)
//	return err, fmt.Sprintf("byte convert to %T failed !", v)
//}

func Byte2Model(data []byte, v anything) error {
	return json.Unmarshal(data, v)
}

func Model2Byte(v anything) ([]byte, error) {
	return json.Marshal(v)
}

func Interface2Model(v interface{}, model *interface{}) error {
	bts, err := Model2Byte(v)
	err = Byte2Model(bts, &model)
	return err
}

func Model2JsonString(v anything) (string, error) {
	bytes, err := Model2Byte(v)
	return string(bytes), err
}

func JsonString2Model(json string, v anything) error {
	return Byte2Model([]byte(json), v)
}

func Byte2Any(data []byte) (interface{}, error) {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
