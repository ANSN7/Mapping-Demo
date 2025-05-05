package main

import (
    "net/http"
    "log"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Create a gauge with two label names ("household_name" and "device").
    electricityConsumption := prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "household_device_electricity_kWh",
            Help: "Electricity consumption by device in households in kWh.",
        },
        // The two label names by which to split the metric.
        []string{"household_name", "device"},
    )

    // Register the gauge with our metrics registry.
    prometheus.MustRegister(electricityConsumption)

    // Set the electricity consumption for different devices in various households.
    electricityConsumption.WithLabelValues("Smiths", "Refrigerator").Set(1.2)
    electricityConsumption.WithLabelValues("Smiths", "Washing Machine").Set(0.7)
    electricityConsumption.WithLabelValues("Martins", "Refrigerator").Set(1.1)
    electricityConsumption.WithLabelValues("Martins", "Oven").Set(0.9)
    electricityConsumption.WithLabelValues("Martins", "Air Conditioner").Set(1.8)
    electricityConsumption.WithLabelValues("Browns", "Refrigerator").Set(1.0)
    electricityConsumption.WithLabelValues("Browns", "Washing Machine").Set(0.6)

    // Expose our custom registry over HTTP on /metrics.
    http.Handle("/metrics", promhttp.Handler())
    log.Fatalln(http.ListenAndServe(":2112", nil))
}