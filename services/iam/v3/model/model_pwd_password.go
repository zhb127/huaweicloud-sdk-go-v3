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
type PwdPassword struct {
	User *PwdPasswordUser `json:"user"`
}

func (o PwdPassword) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PwdPassword", string(data)}, " ")
}
