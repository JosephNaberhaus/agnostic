package model

type Spec struct {
	Name       string            `yaml:"name"`
	Types      []string          `yaml:"types"`
	Properties map[string]string `yaml:"properties"`
}
