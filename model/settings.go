package model

import "strconv"

type Settings struct {
	Verbose                bool
	OS                     string
	MainConfigLocation     string
	ArtifactConfigLocation string
	TargetRepository       string
}

func (a Settings) WithRepositoryNamed(r string) Settings {
	a.TargetRepository = r
	return a
}

func (s Settings) PluginParameters() (params []string) {
	params = append(params, "--verbose="+strconv.FormatBool(s.Verbose))
	if len(s.OS) > 0 {
		params = append(params, "--os="+s.OS)
	}
	if len(s.MainConfigLocation) > 0 {
		params = append(params, "--config="+s.MainConfigLocation)
	}
	if len(s.ArtifactConfigLocation) > 0 {
		params = append(params, "--descriptor="+s.ArtifactConfigLocation)
	}
	if len(s.TargetRepository) > 0 {
		params = append(params, "--repository="+s.TargetRepository)
	}
	return
}
