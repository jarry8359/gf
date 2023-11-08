// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package otelmetric

import (
	"context"
	"go.opentelemetry.io/otel/metric"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gmetric"
)

// localHistogramPerformer is an implementer for interface HistogramPerformer.
type localHistogramPerformer struct {
	config           gmetric.HistogramConfig
	histogram        metric.Float64Histogram
	attributesOption metric.MeasurementOption
}

// newHistogramPerformer creates and returns a HistogramPerformer.
func newHistogramPerformer(meter metric.Meter, config gmetric.HistogramConfig) gmetric.HistogramPerformer {
	histogram, err := meter.Float64Histogram(
		config.Name,
		metric.WithDescription(config.Help),
		metric.WithUnit(config.Unit),
	)
	if err != nil {
		panic(gerror.WrapCodef(
			gcode.CodeInternalError,
			err,
			`create Float64Histogram failed with config: %+v`,
			config,
		))
	}
	return &localHistogramPerformer{
		config:           config,
		histogram:        histogram,
		attributesOption: metric.WithAttributes(attributesToKeyValues(config.Attributes)...),
	}
}

// Record adds a single value to the histogram. The value is usually positive or zero.
func (l *localHistogramPerformer) Record(increment float64, option ...gmetric.Option) {
	l.histogram.Record(
		context.Background(),
		increment,
		l.mergeToRecordOptions(option...)...,
	)
}

func (l *localHistogramPerformer) mergeToRecordOptions(option ...gmetric.Option) []metric.RecordOption {
	var (
		usedOption     gmetric.Option
		observeOptions = []metric.RecordOption{l.attributesOption}
	)
	if len(option) > 0 {
		usedOption = option[0]
	}
	if len(usedOption.Attributes) > 0 {
		observeOptions = append(
			observeOptions,
			metric.WithAttributes(attributesToKeyValues(usedOption.Attributes)...),
		)
	}
	return observeOptions
}