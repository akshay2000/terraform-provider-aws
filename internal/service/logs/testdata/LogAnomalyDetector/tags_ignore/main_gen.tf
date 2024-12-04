# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "aws" {
  default_tags {
    tags = var.provider_tags
  }
  ignore_tags {
    keys = var.ignore_tag_keys
  }
}

resource aws_cloudwatch_log_group "test" {
  count = 2

  name = "${var.rName}-${count.index}"
}

data aws_cloudwatch_log_groups "test" {
  log_group_name_prefix = var.rName

  depends_on = [aws_cloudwatch_log_group.test[0], aws_cloudwatch_log_group.test[1]]
}

resource "aws_logs_log_anomaly_detector" "test" {
  detector_name        = var.rName
  log_arn_group_list   = [aws_cloudwatch_log_group.test.arn]
  evaluation_frequency = "TEN_MIN"
  enabled              = "false"

  tags = var.resource_tags
}
variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}

variable "resource_tags" {
  description = "Tags to set on resource. To specify no tags, set to `null`"
  # Not setting a default, so that this must explicitly be set to `null` to specify no tags
  type     = map(string)
  nullable = true
}

variable "provider_tags" {
  type     = map(string)
  nullable = true
  default  = null
}

variable "ignore_tag_keys" {
  type     = set(string)
  nullable = false
}
