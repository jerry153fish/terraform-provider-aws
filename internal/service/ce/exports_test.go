// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ce

// Exports for use in tests only.
var (
	ResourceAnomalyMonitor      = resourceAnomalyMonitor
	ResourceAnomalySubscription = resourceAnomalySubscription
	ResourceCostAllocationTag   = resourceCostAllocationTag
	ResourceCostCategory        = resourceCostCategory

	FindAnomalyMonitorByARN       = findAnomalyMonitorByARN
	FindAnomalySubscriptionByARN  = findAnomalySubscriptionByARN
	FindCostAllocationTagByTagKey = findCostAllocationTagByTagKey
	FindCostCategoryByARN         = findCostCategoryByARN
)
