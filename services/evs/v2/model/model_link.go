/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 云硬盘uri自描述信息。
type Link struct {
	// 对应的快捷链接。
	Href *string `json:"href,omitempty"`
	// 快捷链接标记名称。
	Rel *string `json:"rel,omitempty"`
}

func (o Link) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Link", string(data)}, " ")
}