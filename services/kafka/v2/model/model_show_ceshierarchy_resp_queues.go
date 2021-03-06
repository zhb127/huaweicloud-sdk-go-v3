/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ShowCeshierarchyRespQueues struct {
	// topic名称。
	Name *string `json:"name,omitempty"`
	// 分区列表。
	Partitions *[]ShowCeshierarchyRespPartitions `json:"partitions,omitempty"`
}

func (o ShowCeshierarchyRespQueues) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCeshierarchyRespQueues", string(data)}, " ")
}
