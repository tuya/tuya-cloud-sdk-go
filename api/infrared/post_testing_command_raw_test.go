package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostTestingCommandRaw(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID  string
		RemoteIndex int
		CategoryID  int
		RawKey      int
	}
	tests := []struct {
		name    string
		args    args
		want    *PostTestingCommandRawResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteIndex: 5129, CategoryID: 5, RawKey: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostTestingCommandRaw(tt.args.InfraredID, tt.args.RemoteIndex, tt.args.CategoryID, tt.args.RawKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostTestingCommandRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
