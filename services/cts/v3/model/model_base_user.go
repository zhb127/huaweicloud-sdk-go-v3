/*
 * cts
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 用户信息。
type BaseUser struct {
	// 账号ID，参见《云审计服务API参考》“获取账号ID和项目ID”章节。
	Id *string `json:"id,omitempty"`
	// 账号名称。
	Name *string `json:"name,omitempty"`
}

func (o BaseUser) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BaseUser", string(data)}, " ")
}