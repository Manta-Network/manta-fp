package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type SymbioticFpMetrics struct {
	// poller metrics
	lastPolledHeight     prometheus.Gauge
	pollerStartingHeight prometheus.Gauge
	// single finality provider metrics
	fpStatus              *prometheus.GaugeVec
	fpLastVotedL1Height   *prometheus.GaugeVec
	fpLastVotedL2Height   *prometheus.GaugeVec
	fpLastProcessedHeight *prometheus.GaugeVec
	fpTotalVotedBlocks    *prometheus.GaugeVec
	fpTotalFailedVotes    *prometheus.CounterVec
	// time keeper
	mu                     sync.Mutex
	previousVoteByFp       map[string]*time.Time
	previousRandomnessByFp map[string]*time.Time
}

// Declare a package-level variable for sync.Once to ensure metrics are registered only once
var symbioticFpMetricsRegisterOnce sync.Once

// Declare a variable to hold the instance of FpMetrics
var symbioticFpMetricsInstance *SymbioticFpMetrics

// NewSymbioticFpMetrics initializes and registers the metrics, using sync.Once to ensure it's done only once
func NewSymbioticFpMetrics() *SymbioticFpMetrics {
	symbioticFpMetricsRegisterOnce.Do(func() {
		symbioticFpMetricsInstance = &SymbioticFpMetrics{
			fpStatus: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: "fp_status",
				Help: "Current status of a finality provider, 0 for active and 1 for Paused",
			}, []string{"operator_address"}),
			lastPolledHeight: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "last_polled_height",
				Help: "The most recent block height checked by the poller",
			}),
			pollerStartingHeight: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "poller_starting_height",
				Help: "The initial block height when the poller started operation",
			}),
			fpLastVotedL1Height: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_last_voted_l1_height",
					Help: "The last l1 block height voted by a finality provider.",
				},
				[]string{"operator_address"},
			),
			fpLastVotedL2Height: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_last_voted_l2_height",
					Help: "The last l2 block height voted by a finality provider.",
				},
				[]string{"operator_address"},
			),
			fpLastProcessedHeight: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_last_processed_height",
					Help: "The last block height processed by a finality provider.",
				},
				[]string{"operator_address"},
			),
			fpTotalVotedBlocks: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_total_voted_blocks",
					Help: "The total number of blocks voted by a finality provider.",
				},
				[]string{"operator_address"},
			),
			fpTotalFailedVotes: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: "fp_total_failed_votes",
					Help: "The total number of failed votes by a finality provider.",
				},
				[]string{"operator_address"},
			),
			mu: sync.Mutex{},
		}

		// Register the metrics with Prometheus
		prometheus.MustRegister(symbioticFpMetricsInstance.fpStatus)
		prometheus.MustRegister(symbioticFpMetricsInstance.lastPolledHeight)
		prometheus.MustRegister(symbioticFpMetricsInstance.pollerStartingHeight)
		prometheus.MustRegister(symbioticFpMetricsInstance.fpLastVotedL1Height)
		prometheus.MustRegister(symbioticFpMetricsInstance.fpLastVotedL2Height)
		prometheus.MustRegister(symbioticFpMetricsInstance.fpLastProcessedHeight)
		prometheus.MustRegister(symbioticFpMetricsInstance.fpTotalVotedBlocks)
		prometheus.MustRegister(symbioticFpMetricsInstance.fpTotalFailedVotes)
	})
	return symbioticFpMetricsInstance
}

// RecordFpStatus records the status of a finality provider
func (fm *SymbioticFpMetrics) RecordFpStatus(operator string, status uint64) {
	fm.fpStatus.WithLabelValues(operator).Set(float64(status))
}

// RecordLastPolledHeight records the most recent block height checked by the poller
func (fm *SymbioticFpMetrics) RecordLastPolledHeight(height uint64) {
	fm.lastPolledHeight.Set(float64(height))
}

// RecordPollerStartingHeight records the initial block height when the poller started operation
func (fm *SymbioticFpMetrics) RecordPollerStartingHeight(height uint64) {
	fm.pollerStartingHeight.Set(float64(height))
}

// RecordFpLastVotedL1Height records the last block height voted by a finality provider
func (fm *SymbioticFpMetrics) RecordFpLastVotedL1Height(operator string, height uint64) {
	fm.fpLastVotedL1Height.WithLabelValues(operator).Set(float64(height))
}

// RecordFpLastVotedL2Height records the last block height voted by a finality provider
func (fm *SymbioticFpMetrics) RecordFpLastVotedL2Height(operator string, height uint64) {
	fm.fpLastVotedL2Height.WithLabelValues(operator).Set(float64(height))
}

// RecordFpLastProcessedHeight records the last block height processed by a finality provider
func (fm *SymbioticFpMetrics) RecordFpLastProcessedHeight(operator string, height uint64) {
	fm.fpLastProcessedHeight.WithLabelValues(operator).Set(float64(height))
}

// IncrementFpTotalVotedBlocks increments the total number of blocks voted by a finality provider
func (fm *SymbioticFpMetrics) IncrementFpTotalVotedBlocks(operator string) {
	fm.fpTotalVotedBlocks.WithLabelValues(operator).Inc()
}

// IncrementFpTotalFailedVotes increments the total number of failed votes by a finality provider
func (fm *SymbioticFpMetrics) IncrementFpTotalFailedVotes(operator string) {
	fm.fpTotalFailedVotes.WithLabelValues(operator).Inc()
}
