package packs

import (
	"fmt"
	"os"
	"strings"

	sh "github.com/codeskyblue/go-sh"
)

// GetBrew returns all software currently installed by Homebrew.
func GetBrew() []Package {
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

// GetBrewCask returns all applications currently installed by Homebrew Cask.
func GetBrewCask() []Package {
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
