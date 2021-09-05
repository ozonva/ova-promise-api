package prometheusmetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/ozonva/ova-promise-api/internal/usecase"
)

type serverMetrics struct {
	createPromiseCounter *prometheus.CounterVec
	updatePromiseCounter *prometheus.CounterVec
	deletePromiseCounter *prometheus.CounterVec
}

const successLabel = "success"

var labelNames = []string{"result"}

func NewServerMetrics() usecase.ServerMetrics {
	return &serverMetrics{
		createPromiseCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "create_promise_count",
			Help: "count of promises created via grpc-server",
		}, labelNames),
		updatePromiseCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "update_promise_count",
			Help: "count of promises updated via grpc-server",
		}, labelNames),
		deletePromiseCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "delete_promise_count",
			Help: "count of promises deleted via grpc-server",
		}, labelNames),
	}
}

func (m *serverMetrics) IncCreatePromiseCounter() {
	m.createPromiseCounter.WithLabelValues(successLabel).Inc()
}

func (m *serverMetrics) IncCreatePromiseCounterByValue(value float64) {
	m.createPromiseCounter.WithLabelValues(successLabel).Add(value)
}

func (m *serverMetrics) IncUpdatePromiseCounter() {
	m.updatePromiseCounter.WithLabelValues(successLabel).Inc()
}

func (m *serverMetrics) IncDeletePromiseCounter() {
	m.deletePromiseCounter.WithLabelValues(successLabel).Inc()
}
