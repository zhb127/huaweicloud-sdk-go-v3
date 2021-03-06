/*
 * RabbitMQ
 *
 * RabbitMQ Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateOrDeleteInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteInstanceTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateOrDeleteInstanceTagResponse", string(data)}, " ")
}
