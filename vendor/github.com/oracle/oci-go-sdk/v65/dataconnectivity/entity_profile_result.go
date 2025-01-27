// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EntityProfileResult A metadata details of a profiling entity result.
type EntityProfileResult struct {

	// Number of columns in the DataFrame (arrow buffer) sent from Java layer. This value is not impacted by the List of attributes to profile as being passed via configuration.
	AttributeCount *int `mandatory:"false" json:"attributeCount"`

	// Number of rows were that were sampled
	SampledRowCount *int `mandatory:"false" json:"sampledRowCount"`

	// The estimated row count in the source.
	EstimatedRowCount *int `mandatory:"false" json:"estimatedRowCount"`
}

func (m EntityProfileResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityProfileResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
