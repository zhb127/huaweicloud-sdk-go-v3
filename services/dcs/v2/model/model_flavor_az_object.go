/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type FlavorAzObject struct {
	// 缓存容量。
	Capacity *string `json:"capacity,omitempty"`
	// 有资源的可用区编码。
	AzCodes *[]string `json:"az_codes,omitempty"`
}

func (o FlavorAzObject) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"FlavorAzObject", string(data)}, " ")
}
