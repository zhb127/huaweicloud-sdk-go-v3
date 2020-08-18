/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
type BackupReplicateReqBody struct {
	// 复制的描述
	Description *string `json:"description,omitempty"`
	// 复制的目标项目ID
	DestinationProjectId string `json:"destination_project_id"`
	// 复制的目标区域
	DestinationRegion string `json:"destination_region"`
	// 复制的目标区域的存储库ID
	DestinationVaultId string `json:"destination_vault_id"`
	// 跨区域复制时，是否启用加速从而缩短复制的时间，如果不指定，默认不启用加速。
	EnableAcceleration *bool `json:"enable_acceleration,omitempty"`
	// 复制名称
	Name *string `json:"name,omitempty"`
}

func (o BackupReplicateReqBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BackupReplicateReqBody", string(data)}, " ")
}