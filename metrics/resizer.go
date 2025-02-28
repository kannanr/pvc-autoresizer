package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	runtimemetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

// Metrics subsystem and all of the keys used by the resizer.
const (
	ResizerSuccessResizeTotalKey = "success_resize_total"
	ResizerFailedResizeTotalKey  = "failed_resize_total"
	ResizerLoopSecondsTotalKey   = "loop_seconds_total"
	ResizerLimitReachedTotalKey  = "limit_reached_total"
)

func init() {
	registerResizerMetrics()
}

type resizerSuccessResizeTotalAdapter struct {
	metric prometheus.CounterVec
}

func (a *resizerSuccessResizeTotalAdapter) Increment(pvcname string, pvcns string) {
	a.metric.With(prometheus.Labels{"persistentvolumeclaim": pvcname, "namespace": pvcns}).Inc()
}

type resizerFailedResizeTotalAdapter struct {
	metric prometheus.CounterVec
}

func (a *resizerFailedResizeTotalAdapter) Increment(pvcname string, pvcns string) {
	a.metric.With(prometheus.Labels{"persistentvolumeclaim": pvcname, "namespace": pvcns}).Inc()
}

type resizerLoopSecondsTotalAdapter struct {
	metric prometheus.Counter
}

func (a *resizerLoopSecondsTotalAdapter) Add(value float64) {
	a.metric.Add(value)
}

type resizerLimitReachedTotalAdapter struct {
	metric prometheus.CounterVec
}

func (a *resizerLimitReachedTotalAdapter) Increment(pvcname string, pvcns string) {
	a.metric.With(prometheus.Labels{"persistentvolumeclaim": pvcname, "namespace": pvcns}).Inc()
}

var (
	resizerSuccessResizeTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: MetricsNamespace,
		Name:      ResizerSuccessResizeTotalKey,
		Help:      "counter that indicates how many volume expansion processing resized succeed.",
	}, []string{"persistentvolumeclaim", "namespace"})

	resizerFailedResizeTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: MetricsNamespace,
		Name:      ResizerFailedResizeTotalKey,
		Help:      "counter that indicates how many volume expansion processing resizes fail.",
	}, []string{"persistentvolumeclaim", "namespace"})

	resizerLoopSecondsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: MetricsNamespace,
		Name:      ResizerLoopSecondsTotalKey,
		Help:      "counter that indicates the sum of seconds spent on volume expansion processing loops.",
	})

	resizerLimitReachedTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: MetricsNamespace,
		Name:      ResizerLimitReachedTotalKey,
		Help:      "counter that indicates how many storage limits were reached.",
	}, []string{"persistentvolumeclaim", "namespace"})

	ResizerSuccessResizeTotal *resizerSuccessResizeTotalAdapter = &resizerSuccessResizeTotalAdapter{metric: *resizerSuccessResizeTotal}
	ResizerFailedResizeTotal  *resizerFailedResizeTotalAdapter  = &resizerFailedResizeTotalAdapter{metric: *resizerFailedResizeTotal}
	ResizerLoopSecondsTotal   *resizerLoopSecondsTotalAdapter   = &resizerLoopSecondsTotalAdapter{metric: resizerLoopSecondsTotal}
	ResizerLimitReachedTotal  *resizerLimitReachedTotalAdapter  = &resizerLimitReachedTotalAdapter{metric: *resizerLimitReachedTotal}
)

func registerResizerMetrics() {
	runtimemetrics.Registry.MustRegister(resizerSuccessResizeTotal)
	runtimemetrics.Registry.MustRegister(resizerFailedResizeTotal)
	runtimemetrics.Registry.MustRegister(resizerLoopSecondsTotal)
	runtimemetrics.Registry.MustRegister(resizerLimitReachedTotal)
}
