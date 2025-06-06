package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// JobsProcessedTotal tracks the total number of processed jobs
	JobsProcessedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "jobs_processed_total",
		Help: "The total number of processed jobs",
	})

	// JobsFailedTotal tracks the total number of failed jobs
	JobsFailedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "jobs_failed_total",
		Help: "The total number of failed jobs",
	})

	// JobDurationSeconds tracks the time taken to process jobs
	JobDurationSeconds = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "job_duration_seconds",
		Help:    "Time taken to process jobs in seconds",
		Buckets: prometheus.DefBuckets,
	})

	// Labeled metrics for per-job-type observability
	JobsProcessedByType = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "jobs_processed_by_type_total",
		Help: "The total number of processed jobs by type",
	}, []string{"type"})

	JobsFailedByType = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "jobs_failed_by_type_total",
		Help: "The total number of failed jobs by type",
	}, []string{"type"})

	JobDurationByType = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "job_duration_by_type_seconds",
		Help:    "Time taken to process jobs by type in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"type"})
)
