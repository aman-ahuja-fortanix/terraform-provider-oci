// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
)

var (
	JmsJmsListJreUsageSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"host_id":          acctest.Representation{RepType: acctest.Optional, Create: `my_host_id_1`},
		"application_id":   acctest.Representation{RepType: acctest.Optional, Create: `my_application_id_1`},
		"application_name": acctest.Representation{RepType: acctest.Optional, Create: `my_appplication_name_1`},
		"time_end":         acctest.Representation{RepType: acctest.Optional, Create: `2021-11-20T01:00:00Z`},
		"time_start":       acctest.Representation{RepType: acctest.Optional, Create: `2021-11-01T01:00:00Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsListJreUsageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsListJreUsageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_jms_list_jre_usage.test_list_jre_usage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_list_jre_usage", "test_list_jre_usage", acctest.Optional, acctest.Create, JmsJmsListJreUsageSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_end", "2021-11-20T01:00:00Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_start", "2021-11-01T01:00:00Z"),

				// JRE usage can be zero
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
