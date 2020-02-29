package service

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// WithMetric use for attach metric to router
func WithMetric(r *mux.Router) func(next http.Handler) http.Handler {
	// 统计请求次数
	HTTPReqTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "go_kit_demo",
		Subsystem: "account",
		Name:      "http_requests_total",
		Help:      "Total number of HTTP requests",
	}, []string{"method", "path"})

	// 统计响应时间
	HTTPReqDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "go_kit_demo",
		Subsystem: "account",
		Name:      "http_request_duration_seconds",
		Help:      "The HTTP request latency in seconds",
		Buckets:   nil,
	}, []string{"method", "path"})

	prometheus.MustRegister(
		HTTPReqTotal,
		HTTPReqDuration,
	)
	// prometheus exporter地址
	r.Path("/metrics").Handler(promhttp.Handler())

	// 返回监控接口中间件
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)

			lvs := []string{r.Method, r.URL.Path}
			HTTPReqTotal.WithLabelValues(lvs...).Inc()
			HTTPReqDuration.WithLabelValues(lvs...).Observe(time.Now().Sub(startTime).Seconds())
		})
	}
}
