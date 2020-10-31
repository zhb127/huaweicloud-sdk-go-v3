/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 创建规则中的监控信息
type MetricInfoForAlarm struct {
	// 指标维度
	Dimensions []MetricsDimension `json:"dimensions"`
	// 指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64。  具体指标名请参见查询指标列表中查询出的指标名。
	MetricName string `json:"metric_name"`
	// 指标命名空间，，例如弹性云服务器命名空间。格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32。说明： 当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为空。
	Namespace string `json:"namespace"`
	// 创建告警规则是选择的资源分组ID，如：rg1603786526428bWbVmk4rP
	ResourceGroupId *string `json:"resource_group_id,omitempty"`
	// 创建告警规则是选择的资源分组名称，如：Resource-Group-ECS-01
	ResourceGroupName *string `json:"resource_group_name,omitempty"`
}

func (o MetricInfoForAlarm) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MetricInfoForAlarm", string(data)}, " ")
}