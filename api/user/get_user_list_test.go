package user

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestGetUsers(t *testing.T) {
	schema := common.Ed.TestDataSchema
	type args struct {
		schema   string
		pageNo   uint
		pageSize uint
	}
	tests := []struct {
		name    string
		args    args
		want    *GetUserListResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{schema: schema, pageNo: 0, pageSize: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserList(tt.args.schema, tt.args.pageNo, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
