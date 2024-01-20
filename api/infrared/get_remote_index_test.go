package infrared

import (
	"testing"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func TestRemoteIndexes(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	categoryID := common.Ed.TestDataCategoryID
	brandId := common.Ed.TestDataBrandID
	type args struct {
		InfraredID string
		CategoryID string
		BrandID    string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetRemoteIndexesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, CategoryID: categoryID, BrandID: brandId},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRemoteIndexes(tt.args.InfraredID, tt.args.CategoryID, tt.args.BrandID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemoteIndexes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}