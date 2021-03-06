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
type UpdateVersionAliasRequest struct {
	FunctionUrn string                         `json:"function_urn"`
	Name        string                         `json:"name"`
	Body        *UpdateVersionAliasRequestBody `json:"body,omitempty"`
}

func (o UpdateVersionAliasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateVersionAliasRequest", string(data)}, " ")
}
