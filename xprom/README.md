# xprom

Gin Prometheus metrics exporter inspired by [github.com/zsais/go-gin-prometheus](https://github.com/zsais/go-gin-prometheus)
Forked from https://github.com/Depado/ginprom

## Install

Simply run:
`add go.mod github.com/vietnguyen87/pkg-golang`

## Differences with go-gin-prometheus

- No support for Prometheus' Push Gateway
- Options on constructor
- Adds a `path` label to get the matched route
- Ability to ignore routes

## Interfaces 

```go 
type Measurer interface {
	GinMiddleware(e *gin.Engine)
}

```

## Usage

```go
package main

import (
	"github.com/vietnguyen87/pkg-golang/xprom"
)

func main() {
	router := gin.New()

	measurer := xprom.New(
		xprom.Namespace("vietnt"),
		xprom.ListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort)),
		xprom.Ignore("/metrics", "/ping", "/swagger/*any"),
		xprom.Subsystem("example_service"),
	)

	//Prometheus Include Recovery
	measurer.GinMiddleware(router)
}
```

### Path

Override the default path (`/metrics`) on which the metrics can be accessed:

```go

func main() {
    router := gin.New()
    
    measurer := xprom.New(
        xprom.Path("/custom/metrics"),
    )
    
    //Prometheus Include Recovery
    measurer.GinMiddleware(router)
}
```

### Namespace

Override the default namespace (`gin`):

```go


func main() {
    router := gin.New()
    
    measurer := xprom.New(
        ginprom.Namespace("vietnt"),
    )
    
    //Prometheus Include Recovery
    measurer.GinMiddleware(router)
}
```

### Subsystem

Override the default (`gonic`) subsystem:

```go
func main() {
    router := gin.New()
    
    measurer := xprom.New(
        ginprom.Subsystem("example_service"),
    )
    
    //Prometheus Include Recovery
    measurer.GinMiddleware(router)
}
```

### Engine

The preferred way to pass the router to xprom:

```go
func main() {
    measurer := xprom.New(
        xprom.Namespace("vietnt"),
        xprom.ListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort)),
        xprom.Ignore("/metrics", "/ping", "/swagger/*any"),
        xprom.Subsystem("example_service"),
    )
    
    //Prometheus Include Recovery
    measurer.GinMiddleware(router)
}
```

### Prometheus Registry

Use a custom `prometheus.Registry` instead of prometheus client's global registry. This option allows
to use ginprom in multiple gin engines in the same process, or if you would like to integrate ginprom with your own
prometheus `Registry`.

```go
registry := prometheus.NewRegistry() // creates new prometheus metric registry
r := gin.New()
p := ginprom.New(
    ginprom.Registry(registry),
)

p.SetListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort))
p.Use(r)
```

### Ignore

Ignore allows to completely ignore some routes. Even though you can apply the
middleware to the only groups you're interested in, it is sometimes useful to
have routes not instrumented.

```go
r := gin.New()
p := ginprom.New(
	ginprom.Engine(r),
	ginprom.Ignore("/api/no/no/no", "/api/super/secret/route")
)

p.SetListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort))
p.Use(r)
```

Note that most of the time this can be solved by gin groups:

```go
r := gin.New()
p := ginprom.New(ginprom.Engine(r))

p.SetListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort))
p.Use(r)
```

### Token

Specify a secret token which Prometheus will use to access the endpoint. If the
token is invalid, the endpoint will return an error.

```go
r := gin.New()
p := ginprom.New(
	ginprom.Engine(r),
	ginprom.Token("supersecrettoken")
)
p.SetListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort))
p.Use(r)
```

### Bucket size

Specify the bucket size for the request duration histogram according to your
expected durations.

```go
r := gin.New()
p := ginprom.New(
	ginprom.Engine(r),
	ginprom.BucketSize([]float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}),
)
p.SetListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort))
p.Use(r)
```

### Summary Objectives 

Specify the bucket size for the request duration histogram according to your
expected durations.

```go
r := gin.New()
p := ginprom.New(
	ginprom.Engine(r),
	ginprom.Objectives( map[float64]float64{
        0.5:  0.01,
        0.95: 0.005,
        0.99: 0.001,
    }),
)
p.SetListenAddress(fmt.Sprintf(":%s", config.GetPrometheusConfig().MetricPort))
p.Use(r)
```

## Troubleshooting

### The instrumentation doesn't seem to work

Make sure you have set the `gin.Engine` in the `ginprom` middleware, either when
initializing it using `ginprom.New(ginprom.Engine(r))` or using the `Use`
function after the initialization.

By design, if the middleware was to panic, it would do so when a route is
called. That's why it just silently fails when no engine has been set.