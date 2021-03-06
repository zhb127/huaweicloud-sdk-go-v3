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
type CreateFileResponse struct {
	// 文件路径。
	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFileResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateFileResponse", string(data)}, " ")
}
