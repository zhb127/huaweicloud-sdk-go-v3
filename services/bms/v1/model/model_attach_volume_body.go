/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 裸金属服务器挂载云硬盘接口请求结构体
type AttachVolumeBody struct {
	VolumeAttachment *VolumeAttachment `json:"volumeAttachment"`
}

func (o AttachVolumeBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AttachVolumeBody", string(data)}, " ")
}
