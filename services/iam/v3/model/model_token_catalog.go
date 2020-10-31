/*
 * IAM
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
type TokenCatalog struct {
	// 该接口所属服务。
	Type string `json:"type"`
	// 服务ID。
	Id string `json:"id"`
	// 服务名称。
	Name string `json:"name"`
	// 终端节点。
	Endpoints []TokenCatalogEndpoint `json:"endpoints"`
}

func (o TokenCatalog) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TokenCatalog", string(data)}, " ")
}