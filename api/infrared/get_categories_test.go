package infrared

import (
"testing"

"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func TestGetCategories(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetCategoriesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCategories(tt.args.InfraredID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}