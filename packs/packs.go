// Copyright © 2018 Clay Dunston <dunstontc@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package packs provides wrappers for package managers.
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
