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
type KeystoneListProjectsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 项目信息列表。
	Projects       *[]ProjectResult `json:"projects,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o KeystoneListProjectsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListProjectsResponse", string(data)}, " ")
}
