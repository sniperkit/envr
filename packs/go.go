package packs

import (
	"fmt"
	"os"
	"path"
	"strings"

	sh "github.com/codeskyblue/go-sh"
)

// GetGo returns all packages installed in your GoPath.
func GetGo() []Package {
	packs := make([]Package, 1)

	goPath := os.Getenv("GOPATH")
	srcPath := path.Join(goPath, "src", "github.com")

	out, err := sh.Command("find", srcPath, "-mindepth", "2", "-maxdepth", "2", "-type", "d").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}

	// searchPath := goPath + "/src/gopkg.in"
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
		dat := strings.Split(line, srcPath+"/")
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
