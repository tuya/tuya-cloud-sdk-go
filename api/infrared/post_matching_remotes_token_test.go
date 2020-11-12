package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostMatchingRemotesToken(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		CategoryID int
		PreToken   string
		Code       string
	}
	tests := []struct {
		name    string
		args    args
		want    *PostMatchingRemotesTokenResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, CategoryID: 5, PreToken: "", Code: "xxxxxxxx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostMatchingRemotesToken(tt.args.InfraredID, tt.args.CategoryID, tt.args.PreToken, tt.args.Code)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostMatchingRemotesToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
