/*
 * CES
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
type CreateResourceGroupResponse struct {
	// 创建的资源分组ID，如：rg1606377637506DmVOENVyL。
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateResourceGroupResponse", string(data)}, " ")
}
