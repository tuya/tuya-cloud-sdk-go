# Tuya Cloud API SDK for Golang

[English](README.md) | [中文版](README_cn.md)

## Introduction

Tuya Cloud API SDK for Golang.

## Get Started

Make sure you have `serverHOST`, `AccessID` and `AccessKey`.

Before running this SDK, you need to initialize it with the following method:

```
config.SetEnv(common.URLCN, "AccessID", "AccessKey")
```

**Example**

If you want to fetch the device info, you can call `device.GetDevice()`:

```
deviceID := "xxx"
got, err := device.GetDevice(deviceID)
if err!=nil{
    xxx
}
// process got
```

## Supported API

|  Method                   | API                                               | Description  |
|  ----                     | ----                                              | ----  |
| token.GetTokenAPI         | GET /v1.0/token?grant_type=1                     | [Get access_token with simple method](https://developer.tuya.com/en/docs/iot/open-api/api-reference/authorization/oauth-management) |
| token.RefreshToken     | GET /v1.0/token/{easy_refresh_token}           | [Refresh token](https://developer.tuya.com/en/docs/iot/open-api/api-reference/authorization/oauth-management) |
| device.GetDevice          | GET /v1.0/devices/{device_id}                  | [Get device details](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/device-management) |
| device.GetDeviceFunctions | GET /v1.0/devices/{deviceId}/functions | [Get function list](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/device-control) |
| device.GetDeviceFunctionByCategory | GET /v1.0/functions/{category} | [Get function list by category](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/device-control) |
| device.GetDeviceStatus | GET /v1.0/devices/{device_id}/status           | [Get device data point details](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/device-control) |
| device.PostDeviceCommand | POST /v1.0/devices/{device_id}/commands       | [Send device command](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/device-control) |
| device.DeleteDevice | DELETE /v1.0/devices/{device_id} | [Remove device](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/device-management) |
| user.PostUserRegister   | POST /v1.0/apps/{schema}/user | [User registration](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-family-management/user-management) |
| user.GetDeviceListByUid | GET /v1.0/users/{uid}/devices | [Get deice list by user ID   ](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/device-management) |
| device.dn.GetDevicesByToken | POST /v1.0/device/paring/token         | [Generate pairing token](https://developer.tuya.com/en/docs/iot/open-api/api-reference/smart-home-devices-management/paring-management) |


## FAQ

### About refreshToken interface

Note: The `refreshToken` interface will return a new `access_token`, even if the old token has not expired.

This logic is already implemented in the `GetToken` method, and generally you do not need to call the `refreshToken` interface.

### Do I need to get the token or refresh the token before calling the API?

No, this logic has been implemented in the API method. The token information will be cached in memory.

### When calling an interface, if the token has expired, do I need to manually call the refresh-token interface?

No, in the `GetToken()` method, it will check whether the token has expired. If token expires, it will be pulled again.

### If a token is refreshed in multiple nodes, then you need to implement `common.TokenLocalManage` interface 

Tuya's cloud token only guarantees that refreshing problems will not occur on the end-users side, but if the token of one user refreshes concurrently on multiple nodes, it will cause one node to be successful while other nodes fail. `GetToken` interface will return an `access_token` and `refresh_token`, but `refresh_token` interface will erase the current `refresh_token` and generate a new token, so the old token will be invalid.

### How to deal with the exception and error of the API method?

If the interface returns an error, it can generally be a URL error or a JSON parsing error, you can contact Tuya technical staff for help.

If error is empty, but the success field of the response is false, you can troubleshoot according to the detailed error information in the `Msg` field.

### Get the device list interface, if there are multiple deviceIDs, how to splice it?

Multiple deviceIDs, separated by commas.

### In the interface for obtaining user list, what does schema refer to?

After creating the App SDK, the channel ID on the detail page is the schema.

### `v1.0/devices/tokens/{{pair_token}}` interface, what does `pair_token` mean? How to get it?

`Pair_token` refers to the network pairing token of an app user, which can be obtained from `v1.0/devices/token`.

### If the API in the SDK is not updated in time, how to implement an API?

There are two ways:

1. You can implement the `common.APIRequest` interface. If it is a POST request, you need to also implement the `RequestBody` interface. Then call `DoAPIRequest()`. For more information, see `UserDevicesAPI`.
2. Submit an issue, and we will handle the update.


## Support

You can get support from Tuya with the following methods:

- Tuya Smart Help Center: [https://support.tuya.com/en/help](https://support.tuya.com/en/help)
- Technical Support: [https://iot.tuya.com/council](https://iot.tuya.com/council)

