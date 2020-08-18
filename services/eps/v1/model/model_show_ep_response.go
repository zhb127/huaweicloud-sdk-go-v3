/*
 * EPS
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
type ShowEpResponse struct {
	EnterpriseProject *EpDetail `json:"enterprise_project,omitempty"`
}

func (o ShowEpResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowEpResponse", string(data)}, " ")
}