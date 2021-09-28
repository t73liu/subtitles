package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "subtitles",
	Short: "CLI for manipulating .srt files",
	Long: `
subtitles is a CLI that can parse .srt files and apply common operations such
as adding/subtracting offset time to all timestamps.
`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
