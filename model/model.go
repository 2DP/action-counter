
package model

import (

)

type Counter struct {
	UUID string `json:"uuid"`
	Count int `json:"count"`
	DurationMillis int64 `json:"duration-millis"`
}


type RedisParam struct {
	Key string `json:"key"`
	Value string `json:"value"`
}