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
type ListSnapshotDataRequest struct {
	PublishDomain *string `json:"publish_domain,omitempty"`
	StartTime     *string `json:"start_time,omitempty"`
	EndTime       *string `json:"end_time,omitempty"`
}

func (o ListSnapshotDataRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSnapshotDataRequest", string(data)}, " ")
}
