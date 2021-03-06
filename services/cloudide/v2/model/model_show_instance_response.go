/*
 * CloudIDE
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {
	Instance *InstancesVo `json:"instance,omitempty"`
	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
