package pairing

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPairingToken(t *testing.T) {
	reqInfo := &PostPairingTokenReq{
		PairingType: "EZ",
		UID:         common.Ed.TestDataUID,
		TimeZoneId:  "Asia/Shanghai",
		HomeID:      "",
	}
	tests := []struct {
		name    string
		reqInfo *PostPairingTokenReq
		want    *PostPairingTokenResponse
		wantErr bool
	}{
		{
			name:    "1",
			reqInfo: reqInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostPairingToken(
				tt.reqInfo.PairingType,
				tt.reqInfo.UID,
				tt.reqInfo.TimeZoneId,
				tt.reqInfo.HomeID,
				"",
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
