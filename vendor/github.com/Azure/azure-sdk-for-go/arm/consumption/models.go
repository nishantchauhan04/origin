package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/shopspring/decimal"
	"net/http"
)

// ErrorDetails is the details of the error.
type ErrorDetails struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// ErrorResponse is error response indicates that the service is not able to process the incoming request. The reason
// is provided in the error message.
type ErrorResponse struct {
	Error *ErrorDetails `json:"error,omitempty"`
}

// MeterDetails is the properties of the meter detail.
type MeterDetails struct {
	MeterName             *string          `json:"meterName,omitempty"`
	MeterCategory         *string          `json:"meterCategory,omitempty"`
	MeterSubCategory      *string          `json:"meterSubCategory,omitempty"`
	Unit                  *string          `json:"unit,omitempty"`
	MeterLocation         *string          `json:"meterLocation,omitempty"`
	TotalIncludedQuantity *decimal.Decimal `json:"totalIncludedQuantity,omitempty"`
	PretaxStandardRate    *decimal.Decimal `json:"pretaxStandardRate,omitempty"`
}

// Operation is a Consumption REST API operation.
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider  *string `json:"provider,omitempty"`
	Resource  *string `json:"resource,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult is result of listing consumption operations. It contains a list of operations and a URL link to
// get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Resource is the Resource model definition.
type Resource struct {
	ID   *string             `json:"id,omitempty"`
	Name *string             `json:"name,omitempty"`
	Type *string             `json:"type,omitempty"`
	Tags *map[string]*string `json:"tags,omitempty"`
}

// UsageDetail is an usage detail resource.
type UsageDetail struct {
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	*UsageDetailProperties `json:"properties,omitempty"`
}

// UsageDetailProperties is the properties of the usage detail.
type UsageDetailProperties struct {
	BillingPeriodID      *string             `json:"billingPeriodId,omitempty"`
	InvoiceID            *string             `json:"invoiceId,omitempty"`
	UsageStart           *date.Time          `json:"usageStart,omitempty"`
	UsageEnd             *date.Time          `json:"usageEnd,omitempty"`
	InstanceName         *string             `json:"instanceName,omitempty"`
	InstanceID           *string             `json:"instanceId,omitempty"`
	InstanceLocation     *string             `json:"instanceLocation,omitempty"`
	Currency             *string             `json:"currency,omitempty"`
	UsageQuantity        *decimal.Decimal    `json:"usageQuantity,omitempty"`
	BillableQuantity     *decimal.Decimal    `json:"billableQuantity,omitempty"`
	PretaxCost           *decimal.Decimal    `json:"pretaxCost,omitempty"`
	IsEstimated          *bool               `json:"isEstimated,omitempty"`
	MeterID              *string             `json:"meterId,omitempty"`
	MeterDetails         *MeterDetails       `json:"meterDetails,omitempty"`
	AdditionalProperties *map[string]*string `json:"additionalProperties,omitempty"`
}

// UsageDetailsListResult is result of listing usage details. It contains a list of available usage details in reverse
// chronological order by billing period.
type UsageDetailsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]UsageDetail `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
}

// UsageDetailsListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client UsageDetailsListResult) UsageDetailsListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}
