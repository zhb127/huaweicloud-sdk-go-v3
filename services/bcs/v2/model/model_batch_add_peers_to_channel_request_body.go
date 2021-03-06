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

// 节点加入通道
type BatchAddPeersToChannelRequestBody struct {
	// 加入某个通道的节点信息
	ChannelPeers *[]BatchAddPeersToChannelRequestBodyChannelPeers `json:"channel_peers,omitempty"`
}

func (o BatchAddPeersToChannelRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchAddPeersToChannelRequestBody", string(data)}, " ")
}
