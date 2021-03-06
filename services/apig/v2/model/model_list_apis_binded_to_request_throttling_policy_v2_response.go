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
type ListApisBindedToRequestThrottlingPolicyV2Response struct {
	// 满足条件的API总数
	Total *int32 `json:"total,omitempty"`
	// 本次返回的API列表长度
	Size *int32 `json:"size,omitempty"`
	// 本次查询返回的API列表
	Apis           *[]ThrottleBindingApiResp `json:"apis,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListApisBindedToRequestThrottlingPolicyV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListApisBindedToRequestThrottlingPolicyV2Response", string(data)}, " ")
}
