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
type CreateInstanceBy3rdResponse struct {
	Result *InstancesResponseInstancesVoResult `json:"result,omitempty"`
	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceBy3rdResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateInstanceBy3rdResponse", string(data)}, " ")
}
