package server

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func UpdateHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := validateUpdateRequest(request); err != nil {
		handleUpdateError(responseWriter, err)
		return
	}

	metricType, metricName, metricValue, err := collectUpdateAttributes(request)

	if err != nil {
		handleUpdateError(responseWriter, err)
		return
	}

	Storage.UpdateMetric(metricType, metricName, metricValue)
	fmt.Fprintf(responseWriter, "Запрос успешно обработан! Новое значение метрики %s: %f", metricName, Storage.metrics[metricType][metricName].GetValue())
	responseWriter.WriteHeader(http.StatusOK)
}

func validateUpdateRequest(request *http.Request) (err error) {
	if request.Method != http.MethodPost {
		err = ErrHTTPMethodNotAllowed
		return
	}

	if request.Header.Get("Content-Type") != "text/plain" {
		err = ErrWrongContextTypeHeader
		return
	}

	return
}

func handleUpdateError(responseWriter http.ResponseWriter, err error) {
	statusCode, ok := StatusCodes[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}
	http.Error(responseWriter, err.Error(), statusCode)
}

func collectUpdateAttributes(request *http.Request) (metricType reflect.Type, metricName string, metricValue interface{}, err error) {
	fullPath := request.URL.Path
	relativePath := strings.TrimPrefix(fullPath, "/update/")
	parts := strings.Split(relativePath, "/")
	metricTypeString, metricName, metricValueString := parts[0], parts[1], parts[2]
	if metricTypeString == "" {
		err = ErrMissedMetricType
		return
	}

	if metricName == "" {
		err = ErrMissedMetricName
		return
	}

	if metricValueString == "" {
		err = ErrMissedMetricValue
		return
	}

	metricType, ok := MetricTypes[metricTypeString]
	if !ok {
		err = ErrUnknownMetricType
		return
	}

	metricValue, err = prepareMetricValue(metricType, metricValueString)
	return
}

func prepareMetricValue(metricType reflect.Type, valueString string) (interface{}, error) {
	switch metricType {
	case reflect.TypeOf(&Gauge{}):
		return strconv.ParseFloat(valueString, 64)
	case reflect.TypeOf(&Counter{}):
		return strconv.ParseInt(valueString, 10, 64)
	default:
		return nil, ErrMissedConversionForMetricType
	}
}
