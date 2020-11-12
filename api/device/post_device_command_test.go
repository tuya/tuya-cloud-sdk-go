package device

import (
	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"

	"testing"
)

func TestPostDeviceCommand(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		deviceID string
		commands []Command
	}
	tests := []struct {
		name    string
		args    args
		want    *PostDeviceCommandResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{deviceID: deviceID, commands: []Command{{Code: "switch_led", Value: false}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostDeviceCommand(tt.args.deviceID, tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendDeviceCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
