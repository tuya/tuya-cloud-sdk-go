package device

import (
	"testing"
)

func TestGetDevicesStatus(t *testing.T) {
	var deviceID = []string{"6ce354901c18c9bcc5lmud"}
	type args struct {
		deviceID []string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetDevicesStatusResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{deviceID: deviceID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDevicesStatus(tt.args.deviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDevices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
