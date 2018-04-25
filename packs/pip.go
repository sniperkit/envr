package packs

import (
	"encoding/json"
	"fmt"
	"os"

	sh "github.com/codeskyblue/go-sh"
)

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
