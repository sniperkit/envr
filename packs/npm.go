package packs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Jeffail/gabs"
	sh "github.com/codeskyblue/go-sh"
)

// NPMPackage struct
type NPMPackage struct {
	Name     string `json:"name"`
	Version  string `json:"version,omitempty"`
	From     string `json:"from,omitempty"`
	Resolved string `json:"resolved,omitempty"`
}

// Implement the Stringer interface for printing
func (p NPMPackage) String() string {
	return fmt.Sprintf("{Package: %s, %s, %s, %s", p.Name, p.From, p.Version, p.Resolved)
}

func getNPM() []NPMPackage {
	packs := make([]NPMPackage, 1)

	out, err := sh.Command("npm", "list", "global", "json", "depth", "0").SetStdin(os.Stdin).Output()
	if err != nil {
		fmt.Println(err)
		defer os.Exit(1)
	}
	jsonParsed, _ := gabs.ParseJSON(out)
	children, _ := jsonParsed.S("dependencies").ChildrenMap()
	for key, child := range children {
		var pack NPMPackage
		err := json.Unmarshal(child.Bytes(), &pack)
		if err != nil {
			log.Fatal(err.Error())
		}
		pack.Name = key
		packs = append(packs, pack)
	}

	json.Unmarshal(out, &packs)

	return packs
}
