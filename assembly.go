package main

type Assembly struct {
	Artifact `yaml:",inline"`
	Api      int         `json:"typhoon-api" yaml:"typhoon-api"`
	Parts    []*Artifact `json:"parts" yaml:"parts"`
}

type Artifact struct {
	Group   string
	Name    string `yaml:"artifact"`
	Version string
	Type    string
}

func (a *Assembly) PostRead() {
	for _, each := range a.Parts {
		if len(each.Group) == 0 {
			each.Group = a.Group
		}
	}
}
