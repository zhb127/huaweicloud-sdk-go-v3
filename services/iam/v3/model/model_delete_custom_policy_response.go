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
type DeleteCustomPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteCustomPolicyResponse", string(data)}, " ")
}
