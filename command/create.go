package command

import (
	"fmt"
	"io/ioutil"
	"strings"

	"encoding/json"
)

// CreateCommand data
type CreateCommand struct {
	Meta
}

// Run Read the .smdl file, build the filesystem, and clone the repositories.
func (c *CreateCommand) Run(args []string) int {
	var m map[string]interface{}
	modelFile := ".\\test.smdl"
	modelData, err := ioutil.ReadFile(modelFile)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("Model Contents: %v", string(modelData))

	json.Unmarshal(modelData, &m)
	var b m.root
	for key := range m {
		fmt.Printf("%v:\n ****************************************", key)
	}

	return 0
}

// Synopsis of the create command
func (c *CreateCommand) Synopsis() string {
	return "Clones repositories described by the .smdl file into the model path."
}

// Help Some information about what will happen when you run create.
func (c *CreateCommand) Help() string {
	helpText := `supermodel create
	The create command searchs in the
	current directory for a file with 
	the .smdl extension, builds namespaces 
	and pulls repositories into context.
`
	return strings.TrimSpace(helpText)
}
