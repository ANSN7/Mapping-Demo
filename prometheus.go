package main

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promauto"
	// "sigs.k8s.io/controller-runtime/pkg/metrics"
	crd "gitlab.eclipse.org/rcarrollred/qos-scheduler/scheduler/api/v1alpha1"
)

var (
	metricPrefix      string
	applicationGroups *prometheus.GaugeVec
	applications      *prometheus.GaugeVec
)

var applicationGroupPhases = [...]crd.ApplicationGroupPhase{
	crd.ApplicationGroupWaiting,
	crd.ApplicationGroupOptimizing,
	crd.ApplicationGroupScheduling,
	crd.ApplicationGroupFailed}

func exportApplicationGroupMetrics(ag crd.ApplicationGroup) {
	fmt.Println(ag.Namespace, ag.Name, ag.Status.Phase)

	// for _, p := range applicationGroupPhases {
	// 	if ag.Status.Phase == p {
			applicationGroups.WithLabelValues("default", "applicationgroup-demo", "Waiting").Set(float64(1))
		// } else {
		// 	applicationGroups.WithLabelValues(ag.Namespace, ag.Name, string(p)).Set(float64(0))
		// }
	// }
}

func main() {
	client, err := api.NewClient(api.Config{
		Address: "http://prometheus-k8s.monitoring.svc.cluster.local:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, "up", time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil && result != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		return
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	// fmt.Printf("Result:\n%v\n", result)

	appGroup := &crd.ApplicationGroup{}


	metricPrefix = "qos_"
	applicationGroups = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricPrefix + "applicationgroups",
		Help: "Application groups with namespace, name, and phase",
	}, []string{"namespace", "name", "phase"})

	// applications = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	// 	Name: metricPrefix + "applications",
	// 	Help: "Applications with namespace, name, and phase",
	// }, []string{"namespace", "name", "phase"})

	prometheus.MustRegister(applicationGroups)

	exportApplicationGroupMetrics(*appGroup)
}
