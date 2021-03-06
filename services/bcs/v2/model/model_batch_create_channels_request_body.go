/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 创建通道的配置
type BatchCreateChannelsRequestBody struct {
	// 通道列表
	Channels *[]ChannelCreateInfo `json:"channels,omitempty"`
}

func (o BatchCreateChannelsRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateChannelsRequestBody", string(data)}, " ")
}
