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
type CreateMemberResponse struct {
	// 请求ID。 注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`
	Member    *Member `json:"member,omitempty"`
}

func (o CreateMemberResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateMemberResponse", string(data)}, " ")
}