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

// This is a auto create Body Object
type UpdateIpGroupRequestBody struct {
	Ipgroup *UpdateIpGroupOption `json:"ipgroup"`
}

func (o UpdateIpGroupRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateIpGroupRequestBody", string(data)}, " ")
}