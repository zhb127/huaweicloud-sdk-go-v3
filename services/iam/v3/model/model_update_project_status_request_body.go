/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
type UpdateProjectStatusRequestBody struct {
	Project *UpdateProjectOption `json:"project"`
}

func (o UpdateProjectStatusRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateProjectStatusRequestBody", string(data)}, " ")
}