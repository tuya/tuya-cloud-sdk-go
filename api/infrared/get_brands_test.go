package infrared

import (
	"testing"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func TestGetBrands(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	categoryID := common.Ed.TestDataCategoryID
	type args struct {
		InfraredID string
		CategoryID string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetBrandsResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, CategoryID: categoryID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBrands(tt.args.InfraredID, tt.args.CategoryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBrands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
