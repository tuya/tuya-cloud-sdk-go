package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPutSetAlias(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		RemoteID   string
		RemoteName string
	}
	tests := []struct {
		name    string
		args    args
		want    *PutSetAliasResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteID: "6cbedcyyyyyyy", RemoteName: "Air conditioning 01"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PutSetAlias(tt.args.InfraredID, tt.args.RemoteID, tt.args.RemoteName)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutSetAlias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
