package infrared

import (
	"testing"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func TestPostAddRemote(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID  string
		CategoryID  string
		BrandID     string
		BrandName   string
		RemoteIndex string
		RemoteName  string
	}
	tests := []struct {
		name    string
		args    args
		want    *PostAddRemoteResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				InfraredID:  deviceID,
				CategoryID:  "5",
				BrandID:     "37",
				BrandName:   "CoolBrand",
				RemoteIndex: "657",
				RemoteName:  "IR control",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostAddRemote(
				tt.args.InfraredID,
				tt.args.CategoryID,
				tt.args.BrandID,
				tt.args.BrandName,
				tt.args.RemoteIndex,
				tt.args.RemoteName,
				"",
				"",
				"",
				"",
				0,
			)
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
