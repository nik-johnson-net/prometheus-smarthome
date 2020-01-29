# Prometheus TP-Link Smarthome Collector

This is a prometheus collector for pulling metrics from a TP-Link Smarthome (Kasa) device. It utilizes
the local protocol exposed by each device rather than the cloud API. Since TP-Links are embedded devices,
this collector is a proxy collector similar to the snmp_exporter tool and functions largely the same way.

## Example

The Collector implements a proxy collector similar to the snmp_exporter for prometheus. In your prometheus.yml,
specify your targets and specify relabel_configs to send the request through the smarthome collector.

```yml
  - job_name: 'prometheus-smarthome'
    static_configs:
      - targets:
        - '192.168.1.40'
        - '192.168.1.41'
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 127.0.0.1:2112  # The prometheus-smarthome's real hostname:port.
```

## Building and running

```sh
cd cmd/prometheus-smarthome
go build
./prometheus-smarthome -port 2112
```

## Contributing

Patches are greatly appreciated, especially to support additional collection points for various device types.

## License

This library is provided under the [MIT License](LICENSE.md)
