package cmdcontainer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pullLong = `Pull a Studios container.`

func PullRun(ident string) (err error) {

	return err
}

var PullCmd = &cobra.Command{

	Use: "pull <name or id>",

	Short: "Pull a Studios container.",

	Long: pullLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Ident'")
			cmd.Usage()
			os.Exit(1)
		}

		var ident string

		if 0 < len(args) {

			ident = args[0]

		}

		err = PullRun(ident)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
