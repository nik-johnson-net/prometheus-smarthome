package main

import (
	"github.com/nik-johnson-net/prometheus-smarthome/pkg"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/nik-johnson-net/prometheus-proxy"
)

func factory(target string) prometheus.Collector {
	return pkg.NewSmarthomeCollector(target)
}

func main() {
	app := proxy.Application{
		CreateFactory: func() proxy.CollectorFactory {
			return factory
		},
	}

	proxy.Main(app)
}