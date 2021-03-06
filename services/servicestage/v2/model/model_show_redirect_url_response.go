/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowRedirectUrlResponse struct {
	// 授权重定向URL。
	Url            *string `json:"url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRedirectUrlResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowRedirectUrlResponse", string(data)}, " ")
}
