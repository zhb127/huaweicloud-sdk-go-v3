/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAppsV2Response struct {
	// 符合条件的APP总数
	Total *int32 `json:"total,omitempty"`
	// 本次查询返回的列表长度
	Size *int32 `json:"size,omitempty"`
	// APP列表
	Apps           *[]AppInfoWithBindNumResp `json:"apps,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAppsV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAppsV2Response", string(data)}, " ")
}
