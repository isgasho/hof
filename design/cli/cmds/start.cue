package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#SetupCommand: schema.#Command & {
	TBD:   " "
	Name:  "setup"
	Usage: "setup"
	Short: "Setup the hof tool"
	Long:  Short
}

#CloneCommand: schema.#Command & {
	TBD:   " "
	Name:  "clone"
	Usage: "clone"
	Short: "Clone a workspace or repository into a new directory"
	Long:  Short
}

#InitCommand: schema.#Command & {
	TBD:   " "
	Name:  "init"
	Usage: "init"
	Short: "Create an empty workspace or initialize an existing directory to one"
	Long:  Short
}
