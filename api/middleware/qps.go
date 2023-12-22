package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var counter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "The total number of HTTP request",
	},
	[]string{"uri"},
)

func init() {
	prometheus.MustRegister(counter)
}

func Qps(c *gin.Context) {
	counter.WithLabelValues(c.Request.RequestURI).Inc()
}
