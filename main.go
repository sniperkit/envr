/**
 * ENVentory provides insights into what you have installed in your current environment.
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cloudogu/spinners"
)

type Package struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	// Path    string `json:"path,omitempty"`
}

type Packages struct {
	Name     string    `json:"name"`
	Packages []Package `json:"packages,omitempty"`
}

func NewPackages(name string, packs []Package) Packages {
	p := Packages{name, packs}
	return p
}

var listPip2 = flag.Bool("pip2", false, "List All Pip2 Packages")
var listPip3 = flag.Bool("pip3", false, "List All Pip3 Packages")
var listGems = flag.Bool("gem", false, "List All Ruby Gems")
var listBrew = flag.Bool("brew", false, "List Homebrew Installs")
var listCask = flag.Bool("brew", false, "List Homebrew Installs")
var listGo = flag.Bool("go", false, "List Installed Go Packages")
var writeList = flag.Bool("write", false, "Write a list of all Packages.")

func main() {
	spinner := spinners.NewDotsSpinner(os.Stdout)
	spinner.Start("taking inventory")

	flag.Parse()

	if *listPip2 {
		for _, p := range getPip2Packages() {
			fmt.Println(p.Name)
		}
	}
	if *listPip3 {
		for _, p := range getPip3Packages() {
			fmt.Println(p.Name)
		}
	}
	if *listGems {
		for _, p := range getRubyGems() {
			fmt.Println(p.Name)
		}
	}
	if *listBrew {
		for _, p := range getBrewPackages() {
			fmt.Println(p.Name)
		}
	}
	if *listCask {
		for _, p := range getCaskPackages() {
			fmt.Println(p.Name)
		}
	}
	if *listGo {
		for _, p := range getGo() {
			fmt.Println(p.Name)
		}
	}

	if *writeList {

		pip2Packs := NewPackages("Pip2", getPip2Packages())
		pip3Packs := NewPackages("Pip3", getPip3Packages())
		brewPacks := NewPackages("Homebrew", getBrewPackages())
		caskPacks := NewPackages("Homebrew Cask", getCaskPackages())
		gemPacks := NewPackages("RubyGems", getRubyGems())
		goPacks := NewPackages("Go", getGo())

		catalog := make([]Packages, 0)
		catalog = append(catalog, pip2Packs, pip3Packs, brewPacks, caskPacks, gemPacks, goPacks)

		var home = os.Getenv("HOME")
		jsonFile, err := os.Create(home + "/" + "inventory-2.json")
		if err != nil {
			fmt.Println("Error creating JSON file:", err.Error())
			// TODO: Handle if file already exists.
			defer os.Exit(1)
		}
		jsonWriter := io.Writer(jsonFile)
		encoder := json.NewEncoder(jsonWriter)
		err = encoder.Encode(&catalog)
		if err != nil {
			fmt.Println("Error encoding JSON to file:", err.Error())
			defer os.Exit(1)
		}

		spinner.Stop()
		fmt.Println("all done :)")
	}

}
