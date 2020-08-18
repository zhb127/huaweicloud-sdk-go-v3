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

//
type AgencyAuth struct {
	Identity *AgencyAuthIdentity `json:"identity"`
}

func (o AgencyAuth) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AgencyAuth", string(data)}, " ")
}