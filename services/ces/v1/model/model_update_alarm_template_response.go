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
type UpdateAlarmTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmTemplateResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateAlarmTemplateResponse", string(data)}, " ")
}
