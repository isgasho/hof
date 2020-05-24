package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
	"github.com/hofstadter-io/hof/lib"
)

var genLong = `  generate all the things, from code to data to config...`

var (
	GenStatsFlag     bool
	GenGeneratorFlag []string
)

func init() {

	GenCmd.Flags().BoolVarP(&GenStatsFlag, "stats", "", false, "Print generator statistics")
	GenCmd.Flags().StringSliceVarP(&GenGeneratorFlag, "generator", "g", nil, "Generators to run, default is all discovered")
}

func GenRun(args []string) (err error) {

	// fmt.Println("GenFlags", GenGeneratorFlag)

	return lib.Gen([]string{}, []string{}, "")

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var GenCmd = &cobra.Command{

	Use: "gen [files...]",

	Aliases: []string{
		"g",
	},

	Short: "generate code, data, and config from your data models and designs",

	Long: genLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = GenRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := GenCmd.HelpFunc()
	usage := GenCmd.UsageFunc()

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
	GenCmd.SetHelpFunc(thelp)
	GenCmd.SetUsageFunc(tusage)

}