package metrics

import (
	"sync"
	"time"

	"github.com/Manta-Network/manta-fp/bbn-fp/proto"
	"github.com/Manta-Network/manta-fp/bbn-fp/store"

	"github.com/prometheus/client_golang/prometheus"
)

type BbnFpMetrics struct {
	// all finality provider metrics
	runningFpGauge prometheus.Gauge
	// poller metrics
	babylonTipHeight     prometheus.Gauge
	lastPolledHeight     prometheus.Gauge
	pollerStartingHeight prometheus.Gauge
	// single finality provider metrics
	fpStatus                        *prometheus.GaugeVec
	fpSecondsSinceLastVote          *prometheus.GaugeVec
	fpSecondsSinceLastRandomness    *prometheus.GaugeVec
	fpLastVotedHeight               *prometheus.GaugeVec
	fpLastProcessedHeight           *prometheus.GaugeVec
	fpLastCommittedRandomnessHeight *prometheus.GaugeVec
	fpTotalBlocksWithoutVotingPower *prometheus.CounterVec
	fpTotalVotedBlocks              *prometheus.GaugeVec
	fpTotalCommittedRandomness      *prometheus.GaugeVec
	fpTotalFailedVotes              *prometheus.CounterVec
	fpTotalFailedRandomness         *prometheus.CounterVec
	// time keeper
	mu                     sync.Mutex
	previousVoteByFp       map[string]*time.Time
	previousRandomnessByFp map[string]*time.Time
}

// Declare a package-level variable for sync.Once to ensure metrics are registered only once
var bbnFpMetricsRegisterOnce sync.Once

// Declare a variable to hold the instance of FpMetrics
var bbNFpMetricsInstance *BbnFpMetrics

// NewBbnFpMetrics initializes and registers the metrics, using sync.Once to ensure it's done only once
func NewBbnFpMetrics() *BbnFpMetrics {
	bbnFpMetricsRegisterOnce.Do(func() {
		bbNFpMetricsInstance = &BbnFpMetrics{
			runningFpGauge: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "total_running_fps",
				Help: "Current number of finality providers that are running",
			}),
			fpStatus: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: "fp_status",
				Help: "Current status of a finality provider",
			}, []string{"fp_btc_pk_hex"}),
			babylonTipHeight: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "babylon_tip_height",
				Help: "The current tip height of the Babylon network",
			}),
			lastPolledHeight: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "last_polled_height",
				Help: "The most recent block height checked by the poller",
			}),
			pollerStartingHeight: prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "poller_starting_height",
				Help: "The initial block height when the poller started operation",
			}),
			fpSecondsSinceLastVote: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_seconds_since_last_vote",
					Help: "Seconds since the last finality sig vote by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpSecondsSinceLastRandomness: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_seconds_since_last_randomness",
					Help: "Seconds since the last public randomness commitment by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpLastVotedHeight: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_last_voted_height",
					Help: "The last block height voted by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpLastProcessedHeight: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_last_processed_height",
					Help: "The last block height processed by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpTotalBlocksWithoutVotingPower: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: "fp_total_blocks_without_voting_power",
					Help: "The total number of blocks without voting power for a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpTotalVotedBlocks: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_total_voted_blocks",
					Help: "The total number of blocks voted by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpTotalCommittedRandomness: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_total_committed_randomness",
					Help: "The total number of randomness commitments by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpLastCommittedRandomnessHeight: prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "fp_last_committed_randomness_height",
					Help: "The last block height with randomness commitment by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpTotalFailedVotes: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: "fp_total_failed_votes",
					Help: "The total number of failed votes by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			fpTotalFailedRandomness: prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: "fp_total_failed_randomness",
					Help: "The total number of failed randomness commitments by a finality provider.",
				},
				[]string{"fp_btc_pk_hex"},
			),
			mu: sync.Mutex{},
		}

		// Register the metrics with Prometheus
		prometheus.MustRegister(bbNFpMetricsInstance.runningFpGauge)
		prometheus.MustRegister(bbNFpMetricsInstance.fpStatus)
		prometheus.MustRegister(bbNFpMetricsInstance.babylonTipHeight)
		prometheus.MustRegister(bbNFpMetricsInstance.lastPolledHeight)
		prometheus.MustRegister(bbNFpMetricsInstance.pollerStartingHeight)
		prometheus.MustRegister(bbNFpMetricsInstance.fpSecondsSinceLastVote)
		prometheus.MustRegister(bbNFpMetricsInstance.fpSecondsSinceLastRandomness)
		prometheus.MustRegister(bbNFpMetricsInstance.fpLastVotedHeight)
		prometheus.MustRegister(bbNFpMetricsInstance.fpLastProcessedHeight)
		prometheus.MustRegister(bbNFpMetricsInstance.fpTotalBlocksWithoutVotingPower)
		prometheus.MustRegister(bbNFpMetricsInstance.fpTotalVotedBlocks)
		prometheus.MustRegister(bbNFpMetricsInstance.fpTotalCommittedRandomness)
		prometheus.MustRegister(bbNFpMetricsInstance.fpLastCommittedRandomnessHeight)
		prometheus.MustRegister(bbNFpMetricsInstance.fpTotalFailedVotes)
		prometheus.MustRegister(bbNFpMetricsInstance.fpTotalFailedRandomness)
	})
	return bbNFpMetricsInstance
}

// DecrementRunningFpGauge decrements the running finality provider gauge
func (fm *BbnFpMetrics) DecrementRunningFpGauge() {
	fm.runningFpGauge.Dec()
}

// IncrementRunningFpGauge increments the running finality provider gauge
func (fm *BbnFpMetrics) IncrementRunningFpGauge() {
	fm.runningFpGauge.Inc()
}

// RecordFpStatus records the status of a finality provider
func (fm *BbnFpMetrics) RecordFpStatus(fpBtcPkHex string, status proto.FinalityProviderStatus) {
	fm.fpStatus.WithLabelValues(fpBtcPkHex).Set(float64(status))
}

// RecordBabylonTipHeight records the current tip height of the Babylon network
func (fm *BbnFpMetrics) RecordBabylonTipHeight(height uint64) {
	fm.babylonTipHeight.Set(float64(height))
}

// RecordLastPolledHeight records the most recent block height checked by the poller
func (fm *BbnFpMetrics) RecordLastPolledHeight(height uint64) {
	fm.lastPolledHeight.Set(float64(height))
}

// RecordPollerStartingHeight records the initial block height when the poller started operation
func (fm *BbnFpMetrics) RecordPollerStartingHeight(height uint64) {
	fm.pollerStartingHeight.Set(float64(height))
}

// RecordFpSecondsSinceLastVote records the seconds since the last finality sig vote by a finality provider
func (fm *BbnFpMetrics) RecordFpSecondsSinceLastVote(fpBtcPkHex string, seconds float64) {
	fm.fpSecondsSinceLastVote.WithLabelValues(fpBtcPkHex).Set(seconds)
}

// RecordFpSecondsSinceLastRandomness records the seconds since the last public randomness commitment by a finality provider
func (fm *BbnFpMetrics) RecordFpSecondsSinceLastRandomness(fpBtcPkHex string, seconds float64) {
	fm.fpSecondsSinceLastRandomness.WithLabelValues(fpBtcPkHex).Set(seconds)
}

// RecordFpLastVotedHeight records the last block height voted by a finality provider
func (fm *BbnFpMetrics) RecordFpLastVotedHeight(fpBtcPkHex string, height uint64) {
	fm.fpLastVotedHeight.WithLabelValues(fpBtcPkHex).Set(float64(height))
}

// RecordFpLastProcessedHeight records the last block height processed by a finality provider
func (fm *BbnFpMetrics) RecordFpLastProcessedHeight(fpBtcPkHex string, height uint64) {
	fm.fpLastProcessedHeight.WithLabelValues(fpBtcPkHex).Set(float64(height))
}

// RecordFpLastCommittedRandomnessHeight record the last height at which a finality provider committed randomness
func (fm *BbnFpMetrics) RecordFpLastCommittedRandomnessHeight(fpBtcPkHex string, height uint64) {
	fm.fpLastCommittedRandomnessHeight.WithLabelValues(fpBtcPkHex).Set(float64(height))
}

// IncrementFpTotalBlocksWithoutVotingPower increments the total number of blocks without voting power for a finality provider
func (fm *BbnFpMetrics) IncrementFpTotalBlocksWithoutVotingPower(fpBtcPkHex string) {
	fm.fpTotalBlocksWithoutVotingPower.WithLabelValues(fpBtcPkHex).Inc()
}

// IncrementFpTotalVotedBlocks increments the total number of blocks voted by a finality provider
func (fm *BbnFpMetrics) IncrementFpTotalVotedBlocks(fpBtcPkHex string) {
	fm.fpTotalVotedBlocks.WithLabelValues(fpBtcPkHex).Inc()
}

// AddToFpTotalVotedBlocks adds a number to the total number of blocks voted by a finality provider
func (fm *BbnFpMetrics) AddToFpTotalVotedBlocks(fpBtcPkHex string, num float64) {
	fm.fpTotalVotedBlocks.WithLabelValues(fpBtcPkHex).Add(num)
}

// AddToFpTotalCommittedRandomness adds a number to the total number of randomness commitments by a finality provider
func (fm *BbnFpMetrics) AddToFpTotalCommittedRandomness(fpBtcPkHex string, num float64) {
	fm.fpTotalCommittedRandomness.WithLabelValues(fpBtcPkHex).Add(num)
}

// IncrementFpTotalFailedVotes increments the total number of failed votes by a finality provider
func (fm *BbnFpMetrics) IncrementFpTotalFailedVotes(fpBtcPkHex string) {
	fm.fpTotalFailedVotes.WithLabelValues(fpBtcPkHex).Inc()
}

// IncrementFpTotalFailedRandomness increments the total number of failed randomness commitments by a finality provider
func (fm *BbnFpMetrics) IncrementFpTotalFailedRandomness(fpBtcPkHex string) {
	fm.fpTotalFailedRandomness.WithLabelValues(fpBtcPkHex).Inc()
}

// RecordFpVoteTime records the time of a finality sig vote by a finality provider
func (fm *BbnFpMetrics) RecordFpVoteTime(fpBtcPkHex string) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	now := time.Now()

	if fm.previousVoteByFp == nil {
		fm.previousVoteByFp = make(map[string]*time.Time)
	}
	fm.previousVoteByFp[fpBtcPkHex] = &now
}

// RecordFpRandomnessTime records the time of a public randomness commitment by a finality provider
func (fm *BbnFpMetrics) RecordFpRandomnessTime(fpBtcPkHex string) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	now := time.Now()

	if fm.previousRandomnessByFp == nil {
		fm.previousRandomnessByFp = make(map[string]*time.Time)
	}
	fm.previousRandomnessByFp[fpBtcPkHex] = &now
}

func (fm *BbnFpMetrics) UpdateFpMetrics(fps []*store.StoredFinalityProvider) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	for _, fp := range fps {
		fm.RecordFpStatus(fp.GetBIP340BTCPK().MarshalHex(), fp.Status)

		if lastVoteTime, ok := fm.previousVoteByFp[fp.GetBIP340BTCPK().MarshalHex()]; ok {
			fm.RecordFpSecondsSinceLastVote(fp.GetBIP340BTCPK().MarshalHex(), time.Since(*lastVoteTime).Seconds())
		}

		if lastRandomnessTime, ok := fm.previousRandomnessByFp[fp.GetBIP340BTCPK().MarshalHex()]; ok {
			fm.RecordFpSecondsSinceLastRandomness(fp.GetBIP340BTCPK().MarshalHex(), time.Since(*lastRandomnessTime).Seconds())
		}
	}
}
