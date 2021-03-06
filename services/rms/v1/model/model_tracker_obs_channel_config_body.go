/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// OBS设置对象
type TrackerObsChannelConfigBody struct {
	// OBS桶名称
	BucketName string `json:"bucket_name"`
	// region id
	RegionId string `json:"region_id"`
}

func (o TrackerObsChannelConfigBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TrackerObsChannelConfigBody", string(data)}, " ")
}
