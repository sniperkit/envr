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
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	// "github.com/cloudogu/spinners"
	p "github.com/dunstontc/envr/packs"
)

var (
	listPip2  = flag.Bool("pip2", false, "List All Pip2 Packages")
	listPip3  = flag.Bool("pip3", false, "List All Pip3 Packages")
	listGems  = flag.Bool("gem", false, "List All Ruby Gems")
	listBrew  = flag.Bool("brew", false, "List Homebrew Installs")
	listNPM   = flag.Bool("npm", false, "List Homebrew Installs")
	listCask  = flag.Bool("cask", false, "List Homebrew Installs")
	listGo    = flag.Bool("go", false, "List Installed Go Packages")
	writeList = flag.Bool("write", false, "Write a list of all Packages.")
)

func main() {
	// cmd.Execute()
	spinner := spinners.NewDotsSpinner(os.Stdout)
	spinner.Start("taking inventory")

	flag.Parse()

	if *listPip2 {
		fmt.Println(p.GetPip2())
	}
	if *listNPM {
		fmt.Println(p.GetNPM())
	}
	if *listPip3 {
		fmt.Println(p.GetPip3())
	}
	if *listGems {
		fmt.Println(p.GetRubyGems())
	}
	if *listBrew {
		fmt.Println(p.GetBrew())
	}
	if *listCask {
		fmt.Println(p.GetBrewCask())
	}
	if *listGo {
		fmt.Println(p.GetGo())
	}

	if *writeList {

		// pip2Packs := p.Packages{"Pip2", p.GetPip2()}
		pip2Packs := p.NewPackages("Pip2", p.GetPip2())
		pip3Packs := p.NewPackages("Pip3", p.GetPip3())
		brewPacks := p.NewPackages("Homebrew", p.GetBrew())
		caskPacks := p.NewPackages("Homebrew Cask", p.GetBrewCask())
		gemPacks := p.NewPackages("RubyGems", p.GetRubyGems())
		goPacks := p.NewPackages("Go", p.GetGo())

		catalog := make([]p.Packages, 0)
		catalog = append(catalog, pip2Packs, pip3Packs, brewPacks, caskPacks, gemPacks, goPacks)

		var home = os.Getenv("HOME")
		jsonFile, err := os.Create(home + "/" + "inventory.json")
		if err != nil {
			fmt.Println("Error creating JSON file:", err.Error())
			// TODO: Handle if file already exists.
			os.Exit(1)
		}
		jsonWriter := io.Writer(jsonFile)
		encoder := json.NewEncoder(jsonWriter)
		err = encoder.Encode(&catalog)
		if err != nil {
			fmt.Println("Error encoding JSON to file:", err.Error())
			os.Exit(1)
		}

	}
	spinner.Stop()

}
