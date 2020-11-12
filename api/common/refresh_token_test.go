package common

import (
	"testing"
)

func TestRefreshTokenAPI(t *testing.T) {
	tests := []struct {
		name    string
		want    *RefreshTokenResponse
		wantErr bool
	}{

		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DoRefreshToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
		})
	}
}
