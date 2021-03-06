/*
 * IAM
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
type KeystoneUpdateMappingResponse struct {
	Mapping        *MappingResult `json:"mapping,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o KeystoneUpdateMappingResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateMappingResponse", string(data)}, " ")
}
