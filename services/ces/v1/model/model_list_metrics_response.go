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

// Response Object
type ListMetricsResponse struct {
	// 指标信息列表
	Metrics        *[]MetricInfoList `json:"metrics,omitempty"`
	MetaData       *MetaData         `json:"meta_data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListMetricsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMetricsResponse", string(data)}, " ")
}
