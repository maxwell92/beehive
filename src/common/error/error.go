package error

import (
	mylog "github.com/maxwell92/gokits/log"
	"encoding/json"
)

var log = mylog.Log

type BeehiveError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewBeehiveError(code int32, data interface{}) *BeehiveError {
	return &BeehiveError{
		Code:    code,
		Message: Errors[code].ErrMsg,
		Data:    data,
	}
}

func (be *BeehiveError) DecodeJson(data string) error {
	err := json.Unmarshal([]byte(data), be)

	if err != nil {
		log.Errorf("BeehiveError DecodeJson Error: err=%s", err)
		return err
	}

	return nil
}

func (be *BeehiveError) EncodeJson() (string, error) {
	data, err := json.Marshal(be)

	if err != nil {
		log.Errorf("BeehiveError EncodeJson Error: err=%s", err)
		return "", err
	}
	return string(data), nil
}

func (be *BeehiveError) String() string {
	errJSON, _ := json.Marshal(be)
	return string(errJSON)
}

func (be *BeehiveError) SetError(code int32, message, data string) {
	be.Code = code
	be.Message = message
	be.Data = data
}

func (be *BeehiveError) GetCode() int32 {
	return be.Code
}

func (be *BeehiveError) GetMessage() string {
	return be.Message
}

func (be *BeehiveError) GetData() interface{} {
	return be.Data
}
