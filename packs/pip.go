package packs

import (
	"encoding/json"
	"fmt"
	"os"

	sh "github.com/codeskyblue/go-sh"
)

// GetPip3 returns all packages globally installed with Pip3.
func GetPip3() []Package {
	packages := make([]Package, 1)
	out, err := sh.Command("pip3", "list", "--form", "json").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	json.Unmarshal(out, &packages)
	return packages
}

// GetPip2 returns all packages globally installed with Pip2.
func GetPip2() []Package {
	packages := make([]Package, 1)
	out, err := sh.Command("pip2", "list", "--form", "json").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	json.Unmarshal(out, &packages)
	return packages
}
