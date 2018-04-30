package packs

import (
	"strings"
)

// Package is comprised of fields from all package types.
type Package struct {
	Name     string `json:"name,omitempty"`     // Package Name
	Source   string `json:"source,omitempty"`   // Source (repo) of a Package
	Version  string `json:"version,omitempty"`  // Package Version
	Path     string `json:"path,omitempty"`     // Path to local installation
	From     string `json:"from,omitempty"`     // NPM
	Resolved string `json:"resolved,omitempty"` // NPM
}

// Packages from a given Package Manager
type Packages struct {
	Name     string    `json:"name"`               // Package Manager
	Packages []Package `json:"packages,omitempty"` // Array of Package
}

// NewPackages news up a new collection of Packages.
func NewPackages(name string, packs []Package) Packages {
	p := Packages{name, packs}
	return p
}

func (p Packages) String() string {
	var str strings.Builder
	for _, r := range p.Packages {
		str.WriteString(r.Name + "\n")
	}
	return str.String()
}

// // PackageManager TODO:
// type PackageManager struct {
// 	Name string
// 	Args []string
// }
