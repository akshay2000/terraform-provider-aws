// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorydb_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccMemoryDBUserDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := "tf-test-" + sdkacctest.RandString(8)
	resourceName := "aws_memorydb_user.test"
	dataSourceName := "data.aws_memorydb_user.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.MemoryDBServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "access_string", resourceName, "access_string"),
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "authentication_mode.0.password_count", resourceName, "authentication_mode.0.password_count"),
					resource.TestCheckResourceAttrPair(dataSourceName, "authentication_mode.0.type", resourceName, "authentication_mode.0.type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "minimum_engine_version", resourceName, "minimum_engine_version"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags.Test", resourceName, "tags.Test"),
					resource.TestCheckResourceAttrPair(dataSourceName, "user_name", resourceName, "user_name"),
				),
			},
		},
	})
}

func TestAccMemoryDBUserDataSource_iam(t *testing.T) {
	ctx := acctest.Context(t)
	rName := "tf-test-" + sdkacctest.RandString(8)
	resourceName := "aws_memorydb_user.test"
	dataSourceName := "data.aws_memorydb_user.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.MemoryDBServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig_iam(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "access_string", resourceName, "access_string"),
					resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "authentication_mode.0.type", resourceName, "authentication_mode.0.type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "authentication_mode.0.password_count", resourceName, "authentication_mode.0.password_count"),
					resource.TestCheckResourceAttrPair(dataSourceName, "minimum_engine_version", resourceName, "minimum_engine_version"),
					resource.TestCheckResourceAttr(dataSourceName, "tags.%", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags.Test", resourceName, "tags.Test"),
					resource.TestCheckResourceAttrPair(dataSourceName, "user_name", resourceName, "user_name"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_memorydb_user" "test" {
  access_string = "on ~* &* +@all"
  user_name     = %[1]q

  authentication_mode {
    type      = "password"
    passwords = ["aaaaaaaaaaaaaaaa"]
  }

  tags = {
    Test = "test"
  }
}

data "aws_memorydb_user" "test" {
  user_name = aws_memorydb_user.test.user_name
}
`, rName)
}

func testAccUserDataSourceConfig_iam(rName string) string {
	return fmt.Sprintf(`
resource "aws_memorydb_user" "test" {
  access_string = "on ~* &* +@all"
  user_name     = %[1]q

  authentication_mode {
    type = "iam"
  }

  tags = {
    Test = "test"
  }
}

data "aws_memorydb_user" "test" {
  user_name = aws_memorydb_user.test.user_name
}
`, rName)
}
