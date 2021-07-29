package device

import (
	"fmt"
	"strings"
	"time"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

const defaultLogSize = 20

type GetDeviceLogsReq struct {
	DeviceID    string
	LogsTypes   []int
	StartRowKey string
	StartTime   int64
	EndTime     int64
	Size        int
}

func (t *GetDeviceLogsReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceLogsReq) API() string {
	return fmt.Sprintf(
		"/v1.0/devices/%s/logs?start_row_key=%s&type=%s&start_time=%d&end_time=%d&size=%d",
		t.DeviceID,
		t.StartRowKey,
		strings.Trim(strings.Replace(fmt.Sprint(t.LogsTypes), " ", ",", -1), "[]"),
		t.StartTime,
		t.EndTime,
		t.Size)
}

// GetDeviceLogs
func GetDeviceLogs(deviceID string, logsTypes []int, startRowKey string, startTime int64, endTime int64, size int) (*GetDeviceLogsResponse, error) {
	a := &GetDeviceLogsReq{DeviceID: deviceID, StartTime: startTime, StartRowKey: startRowKey}
	if logsTypes != nil {
		a.LogsTypes = logsTypes
	} else {
		// All logs
		a.LogsTypes = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	}
	if endTime != 0 {
		a.EndTime = endTime
	} else {
		a.EndTime = time.Now().Unix()
	}
	if size != 0 {
		a.Size = size
	} else {
		a.Size = defaultLogSize
	}
	resp := &GetDeviceLogsResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDeviceLogsResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Logs []struct {
			Code      string `json:"code"`
			Value     string `json:"value"`
			EventTime string `json:"event_time"`
			EventFrom string `json:"event_from"`
			EventID   string `json:"event_id"`
			Status    string `json:"status"`
			Row       string `json:"row"`
		} `json:"logs"`
		HasNext       bool   `json:"has_next"`
		DeviceID      string `json:"device_id"`
		CurrentRowKey string `json:"current_row_key"`
		NextRowKey    string `json:"next_row_key"`
		Count         int64  `json:"count"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
