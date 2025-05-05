package main

import "fmt"

func main(){
	queryTemplate := `sum(rate(container_cpu_usage_seconds_total{pod="%s"}[5m]))`

	query := fmt.Sprintf(queryTemplate, "aaaaaaaaaaaaaa")

	fmt.Print(query)
}