/*
 * CES
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
type UpdateAlarmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateAlarmResponse", string(data)}, " ")
}
