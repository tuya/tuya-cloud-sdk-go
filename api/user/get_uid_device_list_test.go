package user

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func init() {
	common.SetTestEnv()
}

func TestGetDevicesByUID(t *testing.T) {
	UID := common.Ed.TestDataUID
	type args struct {
		uid string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetDeviceListByUIDResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{uid: UID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDeviceListByUID(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
		})
	}
}
