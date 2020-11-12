package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestBrandsByRemoteIndex(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID  string
		CategoryID  int
		RemoteIndex int
		Region      string
		Lang        string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetBrandsByRemoteIndexResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				InfraredID:  deviceID,
				CategoryID:  5,
				RemoteIndex: 9735,
				Region:      "US",
				Lang:        "en",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBrandsByRemoteIndex(tt.args.InfraredID, tt.args.CategoryID, tt.args.RemoteIndex, tt.args.Region, tt.args.Lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBrandsByRemoteIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
