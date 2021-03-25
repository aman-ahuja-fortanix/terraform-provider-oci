// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v37/common"
	"net/http"
)

// GetLogAnalyticsEmBridgeSummaryRequest wrapper for the GetLogAnalyticsEmBridgeSummary operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetLogAnalyticsEmBridgeSummary.go.html to see an example of how to use GetLogAnalyticsEmBridgeSummaryRequest.
type GetLogAnalyticsEmBridgeSummaryRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetLogAnalyticsEmBridgeSummaryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetLogAnalyticsEmBridgeSummaryRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetLogAnalyticsEmBridgeSummaryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetLogAnalyticsEmBridgeSummaryResponse wrapper for the GetLogAnalyticsEmBridgeSummary operation
type GetLogAnalyticsEmBridgeSummaryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LogAnalyticsEmBridgeSummaryReport instance
	LogAnalyticsEmBridgeSummaryReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetLogAnalyticsEmBridgeSummaryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetLogAnalyticsEmBridgeSummaryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
