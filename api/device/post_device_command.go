package device

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type Command struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
}

type PostDeviceCommandReq struct {
	DeviceID string
	Commands []Command
}

func (t *PostDeviceCommandReq) Method() string {
	return common.RequestPost
}

func (t *PostDeviceCommandReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s/commands", t.DeviceID)
}

func (t *PostDeviceCommandReq) Body() []byte {
	m := map[string]interface{}{
		"commands": t.Commands,
	}
	bs, _ := json.Marshal(m)
	return bs
}

// PostDeviceCommand Issue command to device based on device ID
func PostDeviceCommand(deviceID string, commands []Command) (*PostDeviceCommandResponse, error) {
	a := &PostDeviceCommandReq{DeviceID: deviceID, Commands: commands}
	resp := &PostDeviceCommandResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostDeviceCommandResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
