package main

import "github.com/prometheus/client_golang/prometheus"

var (
	gauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Active connections per channel",
		}, []string{"channel_id", "channel_name", "channel_uuid"},
	)
)
