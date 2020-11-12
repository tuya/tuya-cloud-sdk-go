package common

import (
	"fmt"
	"time"
)

func GetTimestamp() string {
	t := fmt.Sprint(time.Now().UnixNano() / 1000000)
	return t
}
