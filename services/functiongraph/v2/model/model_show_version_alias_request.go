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
type ShowVersionAliasRequest struct {
	FunctionUrn string `json:"function_urn"`
	Name        string `json:"name"`
}

func (o ShowVersionAliasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVersionAliasRequest", string(data)}, " ")
}
