// Copyright 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package kstate

import (
	"github.com/wavefronthq/wavefront-collector-for-kubernetes/internal/wf"
)

func buildTags(key, name, ns string, srcTags map[string]string) map[string]string {
	tags := make(map[string]string, len(srcTags)+2)
	tags[key] = name

	if ns != "" {
		tags["namespace_name"] = ns
	}

	for k, v := range srcTags {
		tags[k] = v
	}
	return tags
}

func copyLabels(in map[string]string, out map[string]string) {
	for key, value := range in {
		if len(key) > 0 && len(value) > 0 {
			out["label."+key] = value
		}
	}
}

func copyTags(in map[string]string, out map[string]string) {
	for key, value := range in {
		if len(key) > 0 && len(value) > 0 {
			out[key] = value
		}
	}
}

func metricPoint(prefix, name string, value float64, ts int64, source string, tags map[string]string) *wf.Point {
	return wf.NewPoint(
		prefix+name,
		value,
		ts,
		source,
		tags,
	)
}

func floatVal(i *int32, f float64) float64 {
	if i != nil {
		return float64(*i)
	}
	return f
}
