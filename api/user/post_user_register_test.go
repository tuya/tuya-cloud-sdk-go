package user

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestSignUp(t *testing.T) {
	reqInfo := &PostUserRegisterReq{
		Schema:       "",
		CountryCode:  "86",
		NickName:     "",
		Username:     "",
		Password:     common.Ed.TestDataPassword,
		UsernameType: "",
	}
	tests := []struct {
		name    string
		reqInfo *PostUserRegisterReq
		want    *PostUserRegisterResponse
		wantErr bool
	}{

		{
			name:    "1",
			reqInfo: reqInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostUserRegister(tt.reqInfo)
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
