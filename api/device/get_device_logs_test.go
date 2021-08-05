package device

import (
	"github.com/tuya/tuya-cloud-sdk-go/api/common"

	"testing"
)

func TestGetDeviceLog(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		deviceID    string
		logsTypes   []int
		startRowKey string
		startTime   int64
		endTime     int64
		size        int
	}
	tests := []struct {
		name    string
		args    args
		want    *GetDeviceLogsResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{
				deviceID:    deviceID,
				logsTypes:   nil,
				startRowKey: "",
				startTime:   0,
				endTime:     0,
				size:        0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDeviceLogs(tt.args.deviceID, tt.args.logsTypes, tt.args.startRowKey, tt.args.startTime, tt.args.endTime, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeviceLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
