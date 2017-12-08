package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	PromSpeaker = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "Am_I_Speaker",
			Help: "This represent if speaker or not",
		},
	)

	PromInSession = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "Am_I_Discussing",
			Help: "This represent if discussing or not",
		},
	)

	InboundTCP = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tcp_inbound",
			Help: "The ammount of inbound TCP.",
		},
	)

	OutboundTCP = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tcp_outbound",
			Help: "The ammount of inbound TCP.",
		},
	)

	PromDiscussionParticipants = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "discussion_participants",
			Help: "This represent if discussing participants",
		},
	)

	ENDOutboundTCP = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tcp_outbound_END",
			Help: "The ammount of ends.",
		},
	)

	ABORTOutboundTCP = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tcp_outbound_ABORT",
			Help: "failed aborts.",
		},
	)

	INVOutboundTCP = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "tcp_outbound_INV",
			Help: "The ammount of invites.",
		},
	)




)

func init(){
	prometheus.MustRegister(INVOutboundTCP)
	prometheus.MustRegister(ABORTOutboundTCP)
	prometheus.MustRegister(ENDOutboundTCP)
	prometheus.MustRegister(PromDiscussionParticipants)
	prometheus.MustRegister(OutboundTCP)
	prometheus.MustRegister(InboundTCP)
	prometheus.MustRegister(PromSpeaker)
	prometheus.MustRegister(PromInSession)

}


func ServceMetrics (w http.ResponseWriter, r *http.Request){
	promhttp.Handler().ServeHTTP(w, r)
}