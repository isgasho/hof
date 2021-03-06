package cmdst

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/structural"
)

var pickLong = `pick <what> Cue value(s) from <orig>`

func PickRun(orig string, pick string, entrypoints []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = structural.RunPickFromArgs(orig, pick, entrypoints)

	return err
}

var PickCmd = &cobra.Command{

	Use: "pick <orig> <what> [...entrypoints]",

	Short: "pick <what> Cue value(s) from <orig>",

	Long: pickLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'orig'")
			cmd.Usage()
			os.Exit(1)
		}

		var orig string

		if 0 < len(args) {

			orig = args[0]

		}

		if 1 >= len(args) {
			fmt.Println("missing required argument: 'pick'")
			cmd.Usage()
			os.Exit(1)
		}

		var pick string

		if 1 < len(args) {

			pick = args[1]

		}

		var entrypoints []string

		if 2 < len(args) {

			entrypoints = args[2:]

		}

		err = PickRun(orig, pick, entrypoints)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	ohelp := PickCmd.HelpFunc()
	ousage := PickCmd.UsageFunc()
	help := func(cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		ohelp(cmd, args)
	}
	usage := func(cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		return ousage(cmd)
	}

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	PickCmd.SetHelpFunc(thelp)
	PickCmd.SetUsageFunc(tusage)

}