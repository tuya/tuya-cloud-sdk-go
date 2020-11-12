package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestGetPairingRules(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID  string
		CategoryID  int
		BrandID     int
		RemoteIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *GetPairingRulesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				InfraredID:  deviceID,
				CategoryID:  5,
				BrandID:     27,
				RemoteIndex: 9735,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPairingRules(tt.args.InfraredID, tt.args.CategoryID, tt.args.BrandID, tt.args.RemoteIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPairingRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
