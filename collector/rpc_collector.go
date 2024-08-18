package collector

import (
	"github.com/lbenetti/cisco_exporter/rpc"

	"github.com/lbenetti/client_golang/prometheus"
)

// RPCCollector collects metrics from Cisco using rpc.Client
type RPCCollector interface {
	// Name returns an human readable name for logging and debugging purposes
	Name() string

	// Describe describes the metrics
	Describe(ch chan<- *prometheus.Desc)

	// Collect collects metrics from Cisco
	Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error
}
