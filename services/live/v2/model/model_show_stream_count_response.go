/*
 * Live
 *
 * 数据分析服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowStreamCountResponse struct {
	// 采样数据列表
	DataList       *[]StreamCountData `json:"data_list,omitempty"`
	XRequestId     *string            `json:"X-request-id,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowStreamCountResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowStreamCountResponse", string(data)}, " ")
}
