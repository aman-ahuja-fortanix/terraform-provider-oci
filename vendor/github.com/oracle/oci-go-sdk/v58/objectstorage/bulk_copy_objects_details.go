// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BulkCopyObjectsDetails The parameters required by Object Storage to process a request to copy objects to another bucket.
type BulkCopyObjectsDetails struct {

	// The destination region the object will be copied to, for example "us-ashburn-1".
	DestinationRegion *string `mandatory:"true" json:"destinationRegion"`

	// The destination Object Storage namespace the object will be copied to.
	DestinationNamespace *string `mandatory:"true" json:"destinationNamespace"`

	// The destination bucket the object will be copied to.
	DestinationBucket *string `mandatory:"true" json:"destinationBucket"`

	// List of objects that need to be copied along with their destination names and match conditions.
	ObjectsToCopy []BulkCopyObjectsItem `mandatory:"true" json:"objectsToCopy"`
}

func (m BulkCopyObjectsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkCopyObjectsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}