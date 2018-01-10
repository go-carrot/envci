package envci

import (
	env "github.com/tj/go/env"
)

// Detect current CI
func Detect() *EnvCI {
	variables := &EnvCI{IsCI: false}

	if ci := env.GetDefault("CI", ""); ci != "" {
		variables.IsCI = true
	}

	for service, e := range Services {
		if env.GetDefault(e.Base, "") != "" {
			variables.IsCI = true
			variables.Service = service
			variables.Commit = e.Commit
			variables.Build = e.Build
			variables.Branch = e.Branch
			variables.Job = e.Job
			variables.PR = e.PR
			variables.IsPR = variables.PR != ""
			variables.Slug = e.Slug
			variables.Root = e.Root
			break
		}
	}

	// TODO Get lastet commit
	// if variables.Commit == "" {

	// }

	// TODO Get working branch
	// if variables.Branch == "" {

	// }

	return variables
}

// IsCI tests if CI is in use
func IsCI() bool {
	if env.GetDefault("CI", "") != "" {
		return true
	}

	for _, e := range Services {
		if env.GetDefault(e.Base, "") != "" {
			return true
		}
	}

	return false
}

// EnvCI variables
type EnvCI struct {
	IsCI    bool
	Service string
	Commit  string
	Build   string
	Branch  string
	Job     string
	PR      string
	IsPR    bool
	Slug    string
	Root    string
}
