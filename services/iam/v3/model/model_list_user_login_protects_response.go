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
type ListUserLoginProtectsResponse struct {
	// 登录状态保护信息列表。
	LoginProtects  *[]LoginProtectResult `json:"login_protects,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListUserLoginProtectsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListUserLoginProtectsResponse", string(data)}, " ")
}
