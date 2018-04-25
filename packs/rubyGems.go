package packs

import (
	"fmt"
	"os"
	"strings"

	sh "github.com/codeskyblue/go-sh"
)

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
