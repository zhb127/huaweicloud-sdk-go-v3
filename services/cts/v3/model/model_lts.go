/*
 * cts
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 云日志服务配置
type Lts struct {
	// 是否启用日志服务检索功能。
	IsLtsEnabled *bool `json:"is_lts_enabled,omitempty"`
	// 云审计服务在日志服务中创建的日志组名称。
	LogGroupName *string `json:"log_group_name,omitempty"`
	// 云审计服务在日志服务中创建的日志主题名称。
	LogTopicName *string `json:"log_topic_name,omitempty"`
}

func (o Lts) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Lts", string(data)}, " ")
}