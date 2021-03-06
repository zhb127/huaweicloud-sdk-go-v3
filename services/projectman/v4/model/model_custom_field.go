/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 自定义属性
type CustomField struct {
	// 自定义属性名
	Name *string `json:"name,omitempty"`
	// 自定义属性对应的值
	Value *string `json:"value,omitempty"`
}

func (o CustomField) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CustomField", string(data)}, " ")
}
