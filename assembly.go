package main

type Assembly struct {
	Artifact `yaml:",inline"`
	Api      int        `json:"typhoon-api" yaml:"typhoon-api"`
	Parts    []Artifact `json:"parts" yaml:"parts"`
}

type Artifact struct {
	Group   string
	Name    string `yaml:"artifact"`
	Version string
}
