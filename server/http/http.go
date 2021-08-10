package http

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/realnighthawk/bucky/apm"
	"github.com/realnighthawk/bucky/apm/prometheus"
	"github.com/realnighthawk/bucky/server"
)

const (
	ApplicationRouteGroup = "application"
	MonitoringRouteGroup  = "monitoring"
)

type httpServer struct {
	server.Options
	StartedAt string `json:"started_at,omitempty" yaml:"started_at,omitempty"`
	handler   *handlerGroup
}

type handlerGroup struct {
	application *gin.Engine
	monitoring  *gin.Engine
}

func New(opts server.Options) (*httpServer, error) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	}

	svc := &httpServer{
		Options:   opts,
		StartedAt: time.Now().String(),
	}

	r.GET("/health", healthHandler)
	r.GET("/ping", pingHandler)
	r.GET("/status", svc.statusHandler)

	svc.handler = &handlerGroup{
		application: r,
	}

	return svc, nil
}

func (h *httpServer) Run(runCh chan error) {
	if h.Addresses != nil && len(h.Addresses) > 0 {
		for _, hp := range h.Addresses {
			if hp.Name == ApplicationRouteGroup && h.handler.application != nil {
				go func(hp server.HostPort) {
					err := h.handler.application.Run(fmt.Sprintf("%s:%s", hp.Host, hp.Port))
					if err != nil {
						runCh <- err
					}
				}(hp)
			}
			if hp.Name == MonitoringRouteGroup && h.handler.monitoring != nil {
				go func(hp server.HostPort) {
					err := h.handler.monitoring.Run(fmt.Sprintf("%s:%s", hp.Host, hp.Port))
					if err != nil {
						runCh <- err
					}
				}(hp)
			}
		}
	} else {
		err := h.handler.application.Run("0.0.0.0:80")
		if err != nil {
			runCh <- err
		}
	}
}

func (h *httpServer) EnableMetrics(opts apm.Options) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	}

	if opts.Prometheus.Enabled {
		r.GET("/prometheus", gin.WrapH(prometheus.GetHTTPHandler()))
	}

	h.handler.monitoring = r
}
