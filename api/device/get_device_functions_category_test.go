package device

import (
	"testing"
)

func TestGetDeviceFunctionByCategory(t *testing.T) {
	Category := "kg"
	type args struct {
		Category string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetDeviceFunctionByCategoryResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{Category: Category},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDeviceFunctionByCategory(tt.args.Category)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeviceFunction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
