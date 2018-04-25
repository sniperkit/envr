package packs

// Package is comprised of fields from all package types.
type Package struct {
	Name    string `json:"name,omitempty"`    // Package Name
	Source  string `json:"source,omitempty"`  // Source (repo) of a Package
	Version string `json:"version,omitempty"` // Package Version
	Path    string `json:"path,omitempty"`    // Path to local installation
}

// Packages from a given Package Manager
type Packages struct {
	Name     string    `json:"name"`               // Package Manager
	Packages []Package `json:"packages,omitempty"` // Array of packages
}

// PackageManager TODO:
type PackageManager struct {
	Name string
	Args []string
}

type command []string
