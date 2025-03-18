package server

import (
	"errors"
	"net/http"
)

var (
	ErrMissedMetricType              = errors.New("metric type must be specify")
	ErrMissedMetricName              = errors.New("metric name must be specify")
	ErrUnknownMetricType             = errors.New("metric type is unknown")
	ErrMissedMetricValue             = errors.New("metric value must be specify")
	ErrHTTPMethodNotAllowed          = errors.New("expected POST request")
	ErrWrongContextTypeHeader        = errors.New("allowed only text/plain Content-Type")
	ErrMissedConversionForMetricType = errors.New("metric type is correct, but not convertable yet")

	StatusCodes = map[error]int{
		ErrMissedMetricType:              http.StatusNotFound,
		ErrMissedMetricName:              http.StatusNotFound,
		ErrUnknownMetricType:             http.StatusNotFound,
		ErrMissedMetricValue:             http.StatusBadRequest,
		ErrWrongContextTypeHeader:        http.StatusBadRequest,
		ErrHTTPMethodNotAllowed:          http.StatusMethodNotAllowed,
		ErrMissedConversionForMetricType: http.StatusNotImplemented,
	}
)
