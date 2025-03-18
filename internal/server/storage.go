package server

import (
	"reflect"
	"sync"
)

type MemStorage struct {
	metrics map[reflect.Type]map[string]Metric
	mu      sync.Mutex
}

var (
	Storage = MemStorage{
		metrics: map[reflect.Type]map[string]Metric{
			reflect.TypeOf(&Gauge{}):   map[string]Metric{},
			reflect.TypeOf(&Counter{}): map[string]Metric{},
		},
	}
	MetricTypes = map[string]reflect.Type{
		"gauge":   reflect.TypeOf(&Gauge{}),
		"counter": reflect.TypeOf(&Counter{}),
	}
)

func (storage *MemStorage) UpdateMetric(metricType reflect.Type, name string, value interface{}) {
	metric := storage.FindOrCreateMetric(metricType, name)
	metric.UpdateValue(value)
}

func (storage *MemStorage) FindOrCreateMetric(metricType reflect.Type, name string) Metric {
	metric, ok := storage.metrics[metricType][name]
	if !ok {
		metric := reflect.New(metricType.Elem()).Interface().(Metric)
		metric.SetName(name)
		storage.metrics[metricType][name] = metric
		return metric
	}
	return metric
}
