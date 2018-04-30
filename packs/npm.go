package packs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Jeffail/gabs"
	sh "github.com/codeskyblue/go-sh"
)

// GetNPM returns all globally installed NPM Packages.
func GetNPM() []Package {
	packs := make([]Package, 0)

	out, err := sh.Command("npm", "list", "global", "json", "depth", "0").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	jsonParsed, _ := gabs.ParseJSON(out)
	children, _ := jsonParsed.S("dependencies").ChildrenMap()
	for key, child := range children {
		var pack Package
		err := json.Unmarshal(child.Bytes(), &pack)
		if err != nil {
			log.Fatal(err.Error())
		}
		pack.Name = key
		packs = append(packs, pack)
	}

	// json.Unmarshal(out, &packs)

	return packs
}
