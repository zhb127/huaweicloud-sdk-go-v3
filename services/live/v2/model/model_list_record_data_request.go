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

// Request Object
type ListRecordDataRequest struct {
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
}

func (o ListRecordDataRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRecordDataRequest", string(data)}, " ")
}
