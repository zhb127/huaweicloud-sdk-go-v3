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

// 使用已有CCE集群信息
type CreateRequestBodyCceClusterInfo struct {
	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`
	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`
	// 集群架构类型：X86（VirtualMachine），ARM（ARM64）
	ClusterPlatformType *string `json:"cluster_platform_type,omitempty"`
}

func (o CreateRequestBodyCceClusterInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateRequestBodyCceClusterInfo", string(data)}, " ")
}
