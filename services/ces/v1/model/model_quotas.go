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

//
type Quotas struct {
	// 资源配额列表。
	Resources []Resource `json:"resources"`
}

func (o Quotas) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Quotas", string(data)}, " ")
}
