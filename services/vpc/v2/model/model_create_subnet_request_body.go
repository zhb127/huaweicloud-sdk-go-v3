/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 创建子网对象
type CreateSubnetRequestBody struct {
	Subnet *CreateSubnetOption `json:"subnet"`
}

func (o CreateSubnetRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSubnetRequestBody", string(data)}, " ")
}
