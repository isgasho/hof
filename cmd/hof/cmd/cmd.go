package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/ops"
)

var cmdLong = `run commands from the scripting layer and your _tool.cue files`

func CmdRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = ops.RunCmdFromArgs(args)

	return err
}

var CmdCmd = &cobra.Command{

	Use: "cmd [flags] [cmd] [args]",

	Short: "run commands from the scripting layer and your _tool.cue files",

	Long: cmdLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = CmdRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := CmdCmd.HelpFunc()
	usage := CmdCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/help", "<omit>", 0)
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/usage", "<omit>", 0)
		return usage(cmd)
	}
	CmdCmd.SetHelpFunc(thelp)
	CmdCmd.SetUsageFunc(tusage)

}