package main

import (
	"fmt"
	"github.com/hunjixin/automapper"
	"reflect"
)

type CodecoStatus string

const (
	OK      CodecoStatus = "OK" 
	Warning CodecoStatus = "Warning"
	Error   CodecoStatus = "Error"
)

type ApplicationPhase string

const (
	ApplicationWaiting ApplicationPhase = "Waiting"
	ApplicationScheduling ApplicationPhase = "Scheduling"
	ApplicationPending ApplicationPhase = "Pending"
	ApplicationRunning ApplicationPhase = "Running"
	ApplicationFailing ApplicationPhase = "Failing"
)

type ApplicationCondition struct {
	Entity         string                   `json:"entity,omitempty" protobuf:"bytes,3,opt,name=entity"`
	Reason         string                   `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
}

type CodecoAppStatus struct {
	Status string `json:"status,omitempty" mapping:"Phase"`
	ErrorMsg string `json:"errormsg,omitempty"`
}

type ApplicationStatus struct {
	Phase string `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=ApplicationPhase"`
	Conditions []ApplicationCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"entity" protobuf:"bytes,2,rep,name=conditions"`
}

type Application struct {
	Status ApplicationStatus `json:"status,omitempty"`
}

type CodecoApp struct {
	Status CodecoAppStatus `json:"status,omitempty" mapping:"Status"`
}

func main() {
	user := &CodecoApp{Status: CodecoAppStatus{Status: "OK", ErrorMsg: "hel"}} // Corrected struct initialization
	result := automapper.MustMapper(user, reflect.TypeOf((*Application)(nil)))
	fmt.Println(result)
}