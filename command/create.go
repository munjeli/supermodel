package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateCommand data
type CreateCommand struct {
	Meta
}

// Run Read the .smdl file, build the filesystem, and clone the repositories.
func (c *CreateCommand) Run(args []string) int {
	// Write your code here
	dirname := "." + string(filepath.Separator)
	d, err := os.Open(dirname)
	fmt.Println(d.Readdirnames(-1))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
