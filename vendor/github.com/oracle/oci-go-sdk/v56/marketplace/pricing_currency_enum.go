// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

// PricingCurrencyEnumEnum Enum with underlying type: string
type PricingCurrencyEnumEnum string

// Set of constants representing the allowable values for PricingCurrencyEnumEnum
const (
	PricingCurrencyEnumUsd PricingCurrencyEnumEnum = "USD"
	PricingCurrencyEnumCad PricingCurrencyEnumEnum = "CAD"
	PricingCurrencyEnumInr PricingCurrencyEnumEnum = "INR"
	PricingCurrencyEnumGbp PricingCurrencyEnumEnum = "GBP"
	PricingCurrencyEnumBrl PricingCurrencyEnumEnum = "BRL"
	PricingCurrencyEnumJpy PricingCurrencyEnumEnum = "JPY"
	PricingCurrencyEnumOmr PricingCurrencyEnumEnum = "OMR"
)

var mappingPricingCurrencyEnumEnum = map[string]PricingCurrencyEnumEnum{
	"USD": PricingCurrencyEnumUsd,
	"CAD": PricingCurrencyEnumCad,
	"INR": PricingCurrencyEnumInr,
	"GBP": PricingCurrencyEnumGbp,
	"BRL": PricingCurrencyEnumBrl,
	"JPY": PricingCurrencyEnumJpy,
	"OMR": PricingCurrencyEnumOmr,
}

// GetPricingCurrencyEnumEnumValues Enumerates the set of values for PricingCurrencyEnumEnum
func GetPricingCurrencyEnumEnumValues() []PricingCurrencyEnumEnum {
	values := make([]PricingCurrencyEnumEnum, 0)
	for _, v := range mappingPricingCurrencyEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingCurrencyEnumEnumStringValues Enumerates the set of values in String for PricingCurrencyEnumEnum
func GetPricingCurrencyEnumEnumStringValues() []string {
	return []string{
		"USD",
		"CAD",
		"INR",
		"GBP",
		"BRL",
		"JPY",
		"OMR",
	}
}
