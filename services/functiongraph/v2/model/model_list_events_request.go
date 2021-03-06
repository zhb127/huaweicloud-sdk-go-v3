/*
 * FunctionGraph
 *
 * API v2
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListEventsRequest struct {
	FunctionUrn string `json:"function_urn"`
}

func (o ListEventsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}
