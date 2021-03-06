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

// Request Object
type RemoveProjectRequest struct {
	ProjectId string `json:"project_id"`
}

func (o RemoveProjectRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RemoveProjectRequest", string(data)}, " ")
}
