/*
 * ELB
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
type DeleteL7PolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7PolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteL7PolicyResponse", string(data)}, " ")
}
