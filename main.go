package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type secrets struct {
	Version  string            `yaml:"apiVersion"`
	Kind     string            `yaml:"kind"`
	Metadata metadata          `yaml:"metadata"`
	Type     string            `yaml:"type"`
	Data     map[string]string `yaml:"data"`
}

func New(name string, ns string, data map[string]string) *secrets {
	return &secrets{
		Version: "v1",
		Kind:    "Secret",
		Metadata: metadata{
			Name:      name,
			Namespace: ns,
		},
		Type: "Opaque",
		Data: data,
	}
}

func main() {
	var out, name, namespace string
	flag.StringVar(&out, "o", "", "output file (required)")
	flag.StringVar(&name, "name", "", "secrets name (required)")
	flag.StringVar(&namespace, "ns", "default", "secrets namespace (optional)")
	flag.Parse()
	if out == "" || name == "" || namespace == "" {
		flag.Usage()
		return
	}
	data := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter key (empty to complete): ")
		key, err := reader.ReadString('\n')
		key = strings.TrimSuffix(key, "\n")
		if err != nil {
			log.Fatalf("Failed to read key: %s", err)
		}
		if key == "" {
			break
		}
		fmt.Print("Enter value: ")
		val, err := reader.ReadString('\n')
		val = strings.TrimSuffix(val, "\n")
		b64 := base64.StdEncoding.EncodeToString([]byte(val))
		if err != nil {
			log.Fatalf("Failed to read value: %s", err)
		}
		data[key] = b64
	}
	secrets := New(name, namespace, data)
	d, err := yaml.Marshal(&secrets)
	if err != nil {
		log.Fatalf("Failed to make yaml: %s", err)
	}
	ioutil.WriteFile(out, d, 0644)
	fmt.Printf("Done: %s\n", out)
}
