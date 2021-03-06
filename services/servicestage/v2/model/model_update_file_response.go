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
type UpdateFileResponse struct {
	// 文件路径。
	Path           *string `json:"path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFileResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateFileResponse", string(data)}, " ")
}
