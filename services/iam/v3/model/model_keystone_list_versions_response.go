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
type KeystoneListVersionsResponse struct {
	Versions       *Versions `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o KeystoneListVersionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListVersionsResponse", string(data)}, " ")
}
