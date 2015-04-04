package collectd

import (
	"strings"
	"time"

	"github.com/signalfx/metricproxy/datapoint"
)

// JSONWriteBody is the full POST body of collectd's write_http format
type JSONWriteBody []*JSONWriteFormat

// JSONWriteFormat is the format for collectd json datapoints
type JSONWriteFormat struct {
	Dsnames        []*string  `json:"dsnames"`
	Dstypes        []*string  `json:"dstypes"`
	Host           *string    `json:"host"`
	Interval       *float64   `json:"interval"`
	Plugin         *string    `json:"plugin"`
	PluginInstance *string    `json:"plugin_instance"`
	Time           *float64   `json:"time"`
	TypeS          *string    `json:"type"`
	TypeInstance   *string    `json:"type_instance"`
	Values         []*float64 `json:"values"`
}

func metricTypeFromDsType(dstype *string) datapoint.MetricType {
	if dstype == nil {
		return datapoint.Gauge
	}

	m := map[string]datapoint.MetricType{
		"gauge":    datapoint.Gauge,
		"derive":   datapoint.Counter,
		"counter":  datapoint.Counter,
		"absolute": datapoint.Count,
	}
	v, ok := m[*dstype]
	if ok {
		return v
	}
	return datapoint.Gauge
}

func isNilOrEmpty(str *string) bool {
	return str == nil || *str == ""
}

// NewDatapoint creates a new datapoint from collectd's write_http endpoint JSON format
// defaultDimensions are added to the datapoint created, but will be overridden by any dimension
// values in the JSON
func NewDatapoint(point *JSONWriteFormat, index uint, defaultDimensions map[string]string) *datapoint.Datapoint {
	dstype, val, dsname := point.Dstypes[index], point.Values[index], point.Dsnames[index]
	// if you add another  dimension that we read from the json update this number
	const MaxCollectDDims = 6
	dimensions := make(map[string]string, len(defaultDimensions)+MaxCollectDDims)
	for k, v := range defaultDimensions {
		dimensions[k] = v
	}

	metricType := metricTypeFromDsType(dstype)
	metricName, usedParts := getReasonableMetricName(point, index)
	// Don't add empty dimensions
	addIfNotNullOrEmpty(dimensions, "host", true, point.Host)
	addIfNotNullOrEmpty(dimensions, "plugin", true, point.Plugin)
	addIfNotNullOrEmpty(dimensions, "plugin_instance", true, point.PluginInstance)

	_, usedInMetricName := usedParts["type_instance"]
	addIfNotNullOrEmpty(dimensions, "type_instance", !usedInMetricName, point.TypeInstance)

	_, usedInMetricName = usedParts["type"]
	addIfNotNullOrEmpty(dimensions, "type", !usedInMetricName, point.TypeS)

	_, usedInMetricName = usedParts["dsname"]
	addIfNotNullOrEmpty(dimensions, "dsname", !usedInMetricName, dsname)

	timestamp := time.Unix(0, int64(float64(time.Second)**point.Time))
	return datapoint.New(metricName, dimensions, datapoint.NewFloatValue(*val), metricType, timestamp)
}

func addIfNotNullOrEmpty(dimensions map[string]string, key string, cond bool, val *string) {
	if cond && val != nil && *val != "" {
		dimensions[key] = *val
	}
}

func getReasonableMetricName(point *JSONWriteFormat, index uint) (string, map[string]struct{}) {
	parts := []string{}
	usedParts := make(map[string]struct{})
	if !isNilOrEmpty(point.TypeS) {
		parts = append(parts, *point.TypeS)
		usedParts["type"] = struct{}{}
	}
	if !isNilOrEmpty(point.TypeInstance) {
		parts = append(parts, *point.TypeInstance)
		usedParts["type_instance"] = struct{}{}
	}
	if !isNilOrEmpty(point.Dsnames[index]) && len(point.Dsnames) > 1 {
		parts = append(parts, *point.Dsnames[index])
		usedParts["dsname"] = struct{}{}
	}
	return strings.Join(parts, "."), usedParts
}