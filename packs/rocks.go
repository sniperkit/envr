package packs

import (
	"fmt"
	"os"
	"strings"

	sh "github.com/codeskyblue/go-sh"
)

// - name version, location `[a-zA-Z0-9_-]+(?:\s+)[a-zA-Z0-9.-_]+(?:\s+installed).+`

// GetLuaRocks returns all globally installed rocks.
func GetLuaRocks() []Package {
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
