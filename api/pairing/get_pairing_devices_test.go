package pairing

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPairingDevices(t *testing.T) {
	resp, err := PostPairingToken("EZ", common.Ed.TestDataUID, "Asia/Shanghai", "", "")
	if err != nil {
		t.Errorf("Could not get pairing token, err: %v", err)
		return
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetPairingDevicesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{token: resp.Result.Token},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPairingDevices(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPairingDevices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
