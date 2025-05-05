package main

import (
	"bytes"
	"context"
	// "context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/dynamic"
	// "k8s.io/client-go/tools/clientcmd"
	// "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Masterminds/sprig/v3"
	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"

	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	// "sigs.k8s.io/controller-runtime/pkg/client"
	// "sigs.k8s.io/controller-runtime/pkg/client/config"
)

type RenderData struct {
	Funcs template.FuncMap
	Data  map[string]interface{}
}

func MakeRenderData() RenderData {
	return RenderData{
		Funcs: template.FuncMap{},
		Data:  map[string]interface{}{},
	}
}

// RenderDir will render all manifests in a directory, descending in to subdirectories
// It will perform template substitutions based on the data supplied by the RenderData
func RenderDir(manifestDir string, d *RenderData) ([]*unstructured.Unstructured, error) {
	return RenderDirs([]string{manifestDir}, d)
}

// RenderDirs renders multiple directories, but sorts the discovered files *globally* first.
// In other words, if you have the structure
// - a/001.yaml
// - a/003.yaml
// - b/002.yaml
// It will still render 001, 002, and 003 in order.
func RenderDirs(manifestDirs []string, d *RenderData) ([]*unstructured.Unstructured, error) {
	out := []*unstructured.Unstructured{}

	files := byFilename{}
	for _, dir := range manifestDirs {
		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			// Skip non-manifest files
			if !(strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".json")) {
				return nil
			}

			files = append(files, path)

			return nil
		}); err != nil {
			return nil, errors.Wrap(err, "error listing manifests")
		}
	}
	// sort files by filename, not full path
	sort.Sort(files)

	for _, path := range files {
		objs, err := RenderTemplate(path, d)
		if err != nil {
			return nil, fmt.Errorf("failed to render file %s: %w", path, err)
		}
		out = append(out, objs...)
	}

	return out, nil
}

// sorting boilerplate

type byFilename []string

func (a byFilename) Len() int      { return len(a) }
func (a byFilename) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// sort by filename w/o dir
func (a byFilename) Less(i, j int) bool {
	_, p1 := filepath.Split(a[i])
	_, p2 := filepath.Split(a[j])
	return p1 < p2
}

// RenderTemplate reads, renders, and attempts to parse a yaml or
// json file representing one or more k8s api objects
func RenderTemplate(path string, d *RenderData) ([]*unstructured.Unstructured, error) {
	tmpl := template.New(path).Option("missingkey=error")
	if d.Funcs != nil {
		tmpl.Funcs(d.Funcs)
	}

	// Add universal functions
	tmpl.Funcs(template.FuncMap{"getOr": getOr, "isSet": isSet, "iniEscapeCharacters": iniEscapeCharacters})
	tmpl.Funcs(sprig.TxtFuncMap())

	source, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read manifest %s", path)
	}

	if _, err := tmpl.Parse(string(source)); err != nil {
		return nil, errors.Wrapf(err, "failed to parse manifest %s as template", path)
	}

	rendered := bytes.Buffer{}
	if err := tmpl.Execute(&rendered, d.Data); err != nil {
		return nil, errors.Wrapf(err, "failed to render manifest %s", path)
	}

	out := []*unstructured.Unstructured{}

	// special case - if the entire file is whitespace, skip
	if len(strings.TrimSpace(rendered.String())) == 0 {
		return out, nil
	}

	decoder := yaml.NewYAMLOrJSONDecoder(&rendered, 4096)
	for {
		u := unstructured.Unstructured{}
		if err := decoder.Decode(&u); err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.Wrapf(err, "failed to unmarshal manifest %s", path)
		}
		out = append(out, &u)
	}

	return out, nil
}

func main() {
	manifestDir := "rules"

	// objs := []*uns.Unstructured{}

	// render the manifests on disk
	data := MakeRenderData()
	data.Data["ReleaseVersion"] = os.Getenv("RELEASE_VERSION")
	data.Data["MultiNetworkPolicyImage"] = os.Getenv("MULTUS_NETWORKPOLICY_IMAGE")

	manifests, err := RenderDir(manifestDir, &data)
	if err != nil {
		return
	}
	// objs = append(objs, manifests...)

	// Now 'manifests' contains a list of unstructured Kubernetes objects
	fmt.Printf("Number of manifests rendered: %d\n", len(manifests))

	// Retrieve the default kubeconfig path
		// kubeconfigPath := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
		// if kubeconfigPath == "" {
		// 	fmt.Println("Default kubeconfig path not found.")
		// 	os.Exit(1)
		// }

		// fmt.Println("Kubeconfig path:", kubeconfigPath)

	// 	var kubeconfig = kubeconfigPath
	// config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	//  if err != nil {
	//   fmt.Print("Error building kubeconfig: %v", err)
	//  }

	// // Create a dynamic client
	//  c, err := dynamic.NewForConfig(config)
	//  if err != nil {
	//   fmt.Print("Error creating dynamic client: %v", err)
	//  }else{
	//   fmt.Println("Dynamic client created")
	//  }

	// You can iterate over each manifest and perform operations
	for _, obj := range manifests {
		// Example: Print the Kind of each object
		kind, found, _ := unstructured.NestedString(obj.Object, "kind")
		if !found {
			kind = "<unknown>"
		}
		fmt.Printf("Kind: %s\n", kind)
		fmt.Printf("Kind: %s\n", reflect.TypeOf(obj))
		obj.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "example",
			Version: "monitoring.coreos.com/v1",
			Kind:    "PrometheusRule",
		})

		// result, err := c.Resource(schema.GroupVersionResource{
		// 	Group:   "example",
		// 	Version: "monitoring.coreos.com/v1",
		// 	Resource:    "PrometheusRule",
		//    }).Namespace("default").Create(context.TODO(), obj, metav1.CreateOptions{})

		//   if err != nil {
		//    fmt.Println("Error creating custom resource: %v", err)
		//   }else{
		//   fmt.Println("Custom resource created successfully ", result)
		//   }

		// var c client.Client

		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// c, err = client.New(config.GetConfigOrDie(), client.Options{})
		// cfg := config.GetConfigOrDie()

		// // Create a new Kubernetes client
		// c, err := client.New(cfg, client.Options{})
		// if err != nil {
		// 	fmt.Println("Error creating Kubernetes client:", err)
		// 	return
		// }

		// err = c.Create(context.Background(), obj)
		// if err != nil {
		// 	fmt.Print("Error creating TektonTask!", "ERR", err)
		// } else {
		// 	fmt.Print("Created TektonTask.", "Task")
		// }
	}
}

func getOr(m map[string]interface{}, key string, fallback interface{}) interface{} {
	val, ok := m[key]
	if !ok {
		return fallback
	}

	s, ok := val.(string)
	if ok && s == "" {
		return fallback
	}

	return val
}

// isSet returns the value of m[key] if key exists, otherwise false
// Different from getOr because it will return zero values.
func isSet(m map[string]interface{}, key string) interface{} {
	val, ok := m[key]
	if !ok {
		return false
	}
	return val
}

// iniEscapeCharacters returns the given string with any
// possible reference of '$' escaped.
func iniEscapeCharacters(text string) string {
	return strings.ReplaceAll(text, "$", "\\$")
}
