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

type HandleNotificationOrg struct {
	// 加入的组织
	Name *string `json:"name,omitempty"`
}

func (o HandleNotificationOrg) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"HandleNotificationOrg", string(data)}, " ")
}
