package main

import (
	"encoding/json"
	"fmt"
	"os"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"

	 "sigs.k8s.io/yaml"
)

type EntryConfig struct {
	Version     string `yaml:"version,omitempty"`
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
}

type Entry struct {
	Path   string
	Config *EntryConfig
}

func YamlLoad() {
	r := Entry{Path: "example.yaml"}
	buf, err := os.ReadFile(r.Path)
	if err != nil {
		fmt.Print(err)
	}

    var rule monitoringv1.PrometheusRule
	c := &EntryConfig{}
	err = yaml.Unmarshal(buf, &rule)
	if err != nil {
		fmt.Print(err)
	}
	
	r.Config = c

    fmt.Print(rule)
    s, e:=json.Marshal(rule)
    fmt.Print(string(s))
    fmt.Print(e)
}

func main() {
    YamlLoad()
}