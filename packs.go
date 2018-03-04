package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	sh "github.com/codeskyblue/go-sh"
)

// type NpmData struct {
// 	Version  string `json:"version"`
// 	From     string `json:"from"`
// 	Resolved string `json:"resolved"`
// }
// type NpmPackage struct {
// 	Name string
// 	Data NpmData
// }
// type NpmOutput struct {
// 	deps []NpmPackage `json:"dependencies"`
// }

// func getNPM() []Package {
// 	packs := make([]Package, 1)
// 	out, err := sh.Command("npm", "list", "global", "json", "depth", "0").SetStdin(os.Stdin).Output()
// 	if err != nil {
// 		fmt.Println(err)
// 		defer os.Exit(1)
// 	}
// 	// json.Unmarshal(out, &packages)
// 	gjson.GetBytes(out, dependencies)
// 	return packages
// }

/* json.Unmarshal([]byte, reference to the object to decode the value into) */
func getPip3Packages() []Package {
	packages := make([]Package, 1)
	out, err := sh.Command("pip3", "list", "--form", "json").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	json.Unmarshal(out, &packages)
	return packages
}

func getPip2Packages() []Package {
	packages := make([]Package, 1)
	out, err := sh.Command("pip2", "list", "--form", "json").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	json.Unmarshal(out, &packages)
	return packages
}

func getBrewPackages() []Package {
	packs := make([]Package, 1)
	out, err := sh.Command("brew", "list", "--versions").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	packLines := strings.Split(string(out), "\n")
	for _, line := range packLines {
		dat := strings.SplitAfterN(line, " ", 2)
		if len(line) > 3 {
			pack := Package{
				Name:    dat[0],
				Version: dat[1],
			}
			packs = append(packs, pack)
		}
	}
	return packs
}

func getCaskPackages() []Package {
	packs := make([]Package, 1)
	out, err := sh.Command("brew", "cask", "list", "--versions").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	packLines := strings.Split(string(out), "\n")
	for _, line := range packLines {
		dat := strings.SplitAfterN(line, " ", 2)
		if len(line) > 3 {
			pack := Package{
				Name:    dat[0],
				Version: dat[1],
			}
			packs = append(packs, pack)
		}
	}
	return packs
}

func getRubyGems() []Package {
	packs := make([]Package, 1)
	out, err := sh.Command("gem", "list").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	packLines := strings.Split(string(out), "\n")
	for _, line := range packLines {
		dat := strings.SplitAfterN(line, " ", 2)
		if len(line) > 3 {
			pack := Package{
				Name:    dat[0],
				Version: dat[1],
			}
			packs = append(packs, pack)
		}
	}
	return packs
}

func getGo() []Package {
	packs := make([]Package, 1)

	goPath := os.Getenv("GOPATH")
	srcPath := path.Join(goPath, "src", "github.com")

	out, err := sh.Command("find", srcPath, "-mindepth", "2", "-maxdepth", "2", "-type", "d").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}

	// searchPathTwo := gopath + "/src/gopkg.in"
	// outTwo, errTwo := sh.Command("find", searchPathTwo, "-mindepth", "2", "-maxdepth", "2", "-type", "d").SetStdin(os.Stdin).Output()
	// if errTwo != nil {
	// 	fmt.Println(errTwo)
	// 	defer os.Exit(1)
	// }
	//
	// allPacks := string(out) + "\n" + string(outTwo)

	packLines := strings.Split(string(out), "\n")
	for _, line := range packLines {
		dat := strings.Split(line, searchPath+"/")
		if len(line) >= 1 {
			pack := Package{
				Name: "github.com/" + dat[1],
				// Path: line,
			}
			packs = append(packs, pack)
		}
	}
	return packs
}

func getLuaRocks() []Package {
	packs := make([]Package, 1)
	out, err := sh.Command("luarocks", "list", "porcelain").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	packLines := strings.Split(string(out), "\n")
	for _, line := range packLines {
		if len(line) >= 1 {
			pack := Package{
				Name: line,
				// Path: line,
			}
			packs = append(packs, pack)
		}
	}
	return packs
}
