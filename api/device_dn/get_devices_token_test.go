package device_dn

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func init() {
	common.SetTestEnv()
}

func TestGetDevicesByToken(t *testing.T) {
	PairToken := ""
	type args struct {
		pairToken string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetDevicesByTokenResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{pairToken: PairToken},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDevicesByToken(tt.args.pairToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v\n", got)
		})
	}
}
