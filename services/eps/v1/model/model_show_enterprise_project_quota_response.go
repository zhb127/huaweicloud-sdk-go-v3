/*
 * EPS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowEnterpriseProjectQuotaResponse struct {
	Quotas         *QuotasDetail `json:"quotas,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowEnterpriseProjectQuotaResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowEnterpriseProjectQuotaResponse", string(data)}, " ")
}
