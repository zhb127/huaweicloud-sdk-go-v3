/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteFailedTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFailedTaskResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteFailedTaskResponse", string(data)}, " ")
}
