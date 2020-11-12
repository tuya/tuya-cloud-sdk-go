package common

import (
	"testing"
)

func init() {
	SetTestEnv()
}

func TestGetTokenAPI(t *testing.T) {
	tests := []struct {
		name    string
		want    *GetTokenAPIResponse
		wantErr bool
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTokenAPI()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
