package pkg

import (
	"log"
	"github.com/nik-johnson-net/go-smarthome"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	descEmeterWatts = prometheus.NewDesc("smarthome_emeter_watts", "Amount of watts being used", []string{"deviceID", "deviceAlias", "portID", "portAlias"}, nil)
	descEmeterWattHours = prometheus.NewDesc("smarthome_emeter_watthours", "Amount of watthours consumed", []string{"deviceID", "deviceAlias", "portID", "portAlias"}, nil)
)

type SmarthomeCollector struct {
	target string
}

func NewSmarthomeCollector(target string) *SmarthomeCollector {
	return &SmarthomeCollector{
		target: target,
	}
}

func (s *SmarthomeCollector) Describe(chan<- *prometheus.Desc) {
	return
}

func (s *SmarthomeCollector) Collect(metrics chan<- prometheus.Metric) {
	client := smarthome.NewClient(s.target)

	info, err := client.SysInfo()
	if err != nil {
		log.Printf("failed to collect device %s sysinfo: %s\n", s.target, err.Error())
		return
	}
	
	s.collectDeviceEMeter(metrics, client, info)
}

func (s* SmarthomeCollector) collectDeviceEMeter(metrics chan<- prometheus.Metric, client *smarthome.Client, deviceInfo smarthome.SysInfoResponse) {
	if len(deviceInfo.Children) > 0 {
		for _, child := range deviceInfo.Children {
			emeter, err := client.EMeter(child.ID)
			if err != nil {
				log.Printf("failed to collect device %s child %s emeter: %s\n", deviceInfo.DeviceID, child.ID, err.Error())
				continue
			}
			
			metrics <- prometheus.MustNewConstMetric(
				descEmeterWatts,
				prometheus.GaugeValue,
				emeter.PowerW,
				deviceInfo.DeviceID,
				deviceInfo.Alias,
				child.ID,
				child.Alias,
			)
			metrics <- prometheus.MustNewConstMetric(
				descEmeterWattHours,
				prometheus.CounterValue,
				emeter.TotalWhFloat,
				deviceInfo.DeviceID,
				deviceInfo.Alias,
				child.ID,
				child.Alias,
			)
		}
	} else {
		emeter, err := client.EMeter()
		if err != nil {
			log.Printf("failed to collect device %s emeter: %s\n", deviceInfo.DeviceID, err.Error())
			return
		}

		metrics <- prometheus.MustNewConstMetric(
			descEmeterWatts,
			prometheus.GaugeValue,
			emeter.PowerW,
			deviceInfo.DeviceID,
			deviceInfo.Alias,
			deviceInfo.DeviceID,
			deviceInfo.Alias,
		)
		metrics <- prometheus.MustNewConstMetric(
			descEmeterWattHours,
			prometheus.CounterValue,
			emeter.TotalWhFloat,
			deviceInfo.DeviceID,
			deviceInfo.Alias,
			deviceInfo.DeviceID,
			deviceInfo.Alias,
		)
	}
}