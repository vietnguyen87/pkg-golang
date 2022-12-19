package xprom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	logger "gitlab.marathon.edu.vn/pkg/go/xlogger"
	"net/http"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

const (
	defaultPath                  = "/metrics"
	defaultNs                    = "gin"
	defaultSys                   = "gonic"
	defaultReqCntMetricName      = "requests_total"
	defaultReqPanicCntMetricName = "requests_panic_total"
	defaultReqDurMetricName      = "request_duration"
	defaultReqSzMetricName       = "request_size_bytes"
	defaultResSzMetricName       = "response_size_bytes"
)

var defaultObjectives = map[float64]float64{
	0.9:  0.01,
	0.95: 0.005,
	0.99: 0.001,
}

var defaultBuckets = []float64{0.1, 0.3, 0.5, 1, 2, 3, 5}

// ErrInvalidToken is returned when the provided token is invalid or missing.
var ErrInvalidToken = errors.New("invalid or missing token")

type pmapb struct {
	sync.RWMutex
	values map[string]bool
}

type Measurer interface {
	GinMiddleware(e *gin.Engine)
}

// Prometheus contains the metrics gathered by the instance and its path.
type Prometheus struct {
	reqCnt, reqPanicCnt *prometheus.CounterVec
	//reqDur              *prometheus.SummaryVec
	reqDur *prometheus.HistogramVec
	//reqSz, resSz        prometheus.Summary

	MetricsPath   string
	ListenAddress string
	Namespace     string
	Subsystem     string
	Ignored       pmapb
	Token         string
	Engine        *gin.Engine
	BucketsSize   []float64
	Objectives    map[float64]float64
	Registry      *prometheus.Registry

	RequestPanicCounterMetricName string
	RequestCounterMetricName      string
	RequestDurationMetricName     string
	RequestSizeMetricName         string
	ResponseSizeMetricName        string
}

// Path is an option allowing to set the metrics path when initializing with New.
func Path(path string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.MetricsPath = path
	}
}

// Ignore is used to disable instrumentation on some routes.
func Ignore(paths ...string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.Ignored.Lock()
		defer p.Ignored.Unlock()
		for _, path := range paths {
			p.Ignored.values[path] = true
		}
	}
}

// BucketSize is used to define the default bucket size when initializing with
// New.
func BucketSize(b []float64) func(*Prometheus) {
	return func(p *Prometheus) {
		p.BucketsSize = b
	}
}

// Objectives are used to define the default Objectives when initializing with
// New.
func Objectives(b map[float64]float64) func(*Prometheus) {
	return func(p *Prometheus) {
		p.Objectives = b
	}
}

// Subsystem is an option allowing to set the subsystem when initializing
// with New.
func Subsystem(sub string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.Subsystem = sub
	}
}

// Namespace is an option allowing to set the namespace when initializing
// with New.
func Namespace(ns string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.Namespace = ns
	}
}

// ListenAddress for exposing metrics on address. If not set, it will be exposed at the
// same address of the echo engine that is being used
func ListenAddress(address string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.ListenAddress = address
		if p.ListenAddress != "" {
			p.Engine = gin.Default()
		}
	}
}

// RequestCounterMetricName is an option allowing to set the request counter metric name.
func RequestCounterMetricName(reqCntMetricName string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.RequestCounterMetricName = reqCntMetricName
	}
}

// RequestPanicCounterMetricName is an option allowing to set the request counter metric name.
func RequestPanicCounterMetricName(reqPanicCntMetricName string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.RequestPanicCounterMetricName = reqPanicCntMetricName
	}
}

// RequestDurationMetricName is an option allowing to set the request duration metric name.
func RequestDurationMetricName(reqDurMetricName string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.RequestDurationMetricName = reqDurMetricName
	}
}

// RequestSizeMetricName is an option allowing to set the request size metric name.
func RequestSizeMetricName(reqSzMetricName string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.RequestSizeMetricName = reqSzMetricName
	}
}

// ResponseSizeMetricName is an option allowing to set the response size metric name.
func ResponseSizeMetricName(resDurMetricName string) func(*Prometheus) {
	return func(p *Prometheus) {
		p.ResponseSizeMetricName = resDurMetricName
	}
}

// Engine is an option allowing to set the gin engine when initializing with New.
// Example:
// r := gin.Default()
// p := ginprom.New(Engine(r))
func Engine(e *gin.Engine) func(*Prometheus) {
	return func(p *Prometheus) {
		p.Engine = e
	}
}

// Registry is an option allowing to set a  *prometheus.Registry with New.
// Use this option if you want to use a custom Registry instead of a global one that prometheus
// client uses by default
// Example:
// r := gin.Default()
// p := ginprom.New(Registry(r))
func Registry(r *prometheus.Registry) func(*Prometheus) {
	return func(p *Prometheus) {
		p.Registry = r
	}
}

// New will initialize a new Prometheus instance with the given options.
// If no options are passed, sane defaults are used.
// If a router is passed using the Engine() option, this instance will
// automatically bind to it.
func New(options ...func(*Prometheus)) *Prometheus {
	p := &Prometheus{
		MetricsPath:                   defaultPath,
		Namespace:                     defaultNs,
		Subsystem:                     defaultSys,
		BucketsSize:                   defaultBuckets,
		Objectives:                    defaultObjectives,
		RequestPanicCounterMetricName: defaultReqPanicCntMetricName,
		RequestCounterMetricName:      defaultReqCntMetricName,
		RequestDurationMetricName:     defaultReqDurMetricName,
		RequestSizeMetricName:         defaultReqSzMetricName,
		ResponseSizeMetricName:        defaultResSzMetricName,
	}
	//p.customGauges.values = make(map[string]prometheus.GaugeVec)
	p.Ignored.values = make(map[string]bool)
	for _, option := range options {
		option(p)
	}

	p.register()
	if p.Engine != nil {
		registerer, gatherer := p.getRegistererAndGatherer()
		p.Engine.GET(p.MetricsPath, prometheusHandler(p.Token, registerer, gatherer))
	}

	p.runServer()

	return p
}

func (p *Prometheus) getRegistererAndGatherer() (prometheus.Registerer, prometheus.Gatherer) {
	if p.Registry == nil {
		return prometheus.DefaultRegisterer, prometheus.DefaultGatherer
	}
	return p.Registry, p.Registry
}

func (p *Prometheus) register() {
	registerer, _ := p.getRegistererAndGatherer()
	p.reqCnt = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.Namespace,
			Subsystem: p.Subsystem,
			Name:      p.RequestCounterMetricName,
			Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
		}, []string{"code", "method", "host", "path"},
	)
	registerer.MustRegister(p.reqCnt)

	p.reqPanicCnt = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: p.Namespace,
			Subsystem: p.Subsystem,
			Name:      p.RequestPanicCounterMetricName,
			Help:      "How many HTTP requests get panic.",
		},
		[]string{"service", "host", "path"},
	)
	registerer.MustRegister(p.reqPanicCnt)

	//p.reqDur = prometheus.NewSummaryVec(prometheus.SummaryOpts{
	//	Namespace:  p.Namespace,
	//	Subsystem:  p.Subsystem,
	//	Objectives: p.Objectives,
	//	Name:       p.RequestDurationMetricName,
	//	Help:       "The HTTP request latency objectives.",
	//}, []string{"method", "path", "host"})
	//
	//registerer.MustRegister(p.reqDur)

	p.reqDur = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: p.Namespace,
		Subsystem: p.Subsystem,
		Buckets:   p.BucketsSize,
		Name:      p.RequestDurationMetricName,
		Help:      "The HTTP request latency bucket.",
	}, []string{"method", "path", "host"})

	registerer.MustRegister(p.reqDur)

	/*p.reqSz = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: p.Namespace,
			Subsystem: p.Subsystem,
			Name:      p.RequestSizeMetricName,
			Help:      "The HTTP request sizes in bytes.",
		},
	)
	registerer.MustRegister(p.reqSz)

	p.resSz = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: p.Namespace,
			Subsystem: p.Subsystem,
			Name:      p.ResponseSizeMetricName,
			Help:      "The HTTP response sizes in bytes.",
		},
	)
	registerer.MustRegister(p.resSz)*/
}

func (p *Prometheus) isIgnored(path string) bool {
	p.Ignored.RLock()
	defer p.Ignored.RUnlock()
	_, ok := p.Ignored.values[path]
	return ok
}

// GinMiddleware is a method that should be used if the engine is set after middleware
// initialization.
func (p *Prometheus) GinMiddleware(e *gin.Engine) {
	registerer, gatherer := p.getRegistererAndGatherer()
	e.Use(p.HandlerFunc())
	p.Engine = e
	p.Engine.GET(p.MetricsPath, prometheusHandler(p.Token, registerer, gatherer))
}

func (p *Prometheus) runServer() {
	if p.ListenAddress != "" {
		go p.Engine.Run(p.ListenAddress)
	}
}

// HandlerFunc is a gin middleware that can be used to generate metrics for a
// single handler
func (p *Prometheus) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()

		defer func() {
			if v := recover(); v != nil {
				log := logger.CToL(c.Request.Context(), p.Subsystem)
				p.reqPanicCnt.WithLabelValues(p.Subsystem, c.Request.Host, path).Inc()
				p.reqCnt.WithLabelValues(fmt.Sprint(http.StatusInternalServerError), c.Request.Method, c.Request.Host, path).Inc()
				log.Errorf("[Recovery] panic method %v, path %v, err %v, trace %v",
					c.Request.Method,
					c.Request.URL.EscapedPath(),
					v,
					string(debug.Stack()))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		if path == "" || p.isIgnored(path) {
			c.Next()
			return
		}

		//reqSz := computeApproximateRequestSize(c.Request)
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)
		//resSz := float64(c.Writer.Size())

		p.reqCnt.WithLabelValues(status, c.Request.Method, c.Request.Host, path).Inc()
		p.reqDur.WithLabelValues(c.Request.Method, path, c.Request.Host).Observe(elapsed)

		//p.reqSz.Observe(float64(reqSz))
		//p.resSz.Observe(resSz)
	}
}

func prometheusHandler(token string, registerer prometheus.Registerer, gatherer prometheus.Gatherer) gin.HandlerFunc {
	h := promhttp.InstrumentMetricHandler(
		registerer, promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{}),
	)
	return func(c *gin.Context) {
		if token == "" {
			h.ServeHTTP(c.Writer, c.Request)
			return
		}

		header := c.Request.Header.Get("Authorization")
		if header == "" {
			c.String(http.StatusUnauthorized, ErrInvalidToken.Error())
			return
		}

		bearer := fmt.Sprintf("Bearer %s", token)
		if header != bearer {
			c.String(http.StatusUnauthorized, ErrInvalidToken.Error())
			return
		}

		h.ServeHTTP(c.Writer, c.Request)
	}
}
