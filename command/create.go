package command

// algorithm for building the model
// list the keys in the map
// start with repos and fall thru actually this is wrong
// if key == repos:
// 	os.getcwd() is path to clone url
//  (make repo struct though for reads later on branches)
// else:
//   if namespace in key.path:
//		namespace = namespace + key
// 	mkdirs(keys)
//

//
import (
	"fmt"
	"os"
	"strings"

	"github.com/munjeli/modeler"
)

// CreateCommand data
type CreateCommand struct {
	Meta
}

// Run Read the .smdl file, build the filesystem, and clone the repositories.
func (c *CreateCommand) Run(args []string) int {
	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("boo")
	}
	modelFile := workingDirectory + "\\test.smdl"
	modeler.LoadModel(modelFile)

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
