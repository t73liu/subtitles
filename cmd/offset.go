package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/t73liu/subtitles/srt"
)

// offsetCmd represents the offset command
var offsetCmd = &cobra.Command{
	Use:   "offset file [flags]",
	Short: "Applies offset to all timestamps",
	Long: `
Command applies offset to all timestamps. This is useful when fixing subtitles
that are not properly synced with video.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("provide one file path")
		}

		file := args[0]
		fileInfo, err := os.Stat(file)
		if err != nil {
			return fmt.Errorf("failed to get file info: %s", file)
		}

		if fileInfo.IsDir() {
			return fmt.Errorf("provide file path instead of directory: %s", file)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		durationStr, err := cmd.Flags().GetString("duration")
		if err != nil || durationStr == "" {
			return fmt.Errorf("could not get --duration argument: %s", err)
		}

		outputFile, err := cmd.Flags().GetString("output")
		if err != nil || outputFile == "" {
			return fmt.Errorf("could not get --output argument: %s", err)
		}

		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			return fmt.Errorf("could not parse --duration argument: %s", err)
		}

		subtitles, err := srt.ReadSRTFile(args[0])
		if err != nil {
			return fmt.Errorf("could not read .srt file: %s", err)
		}

		for _, sub := range subtitles {
			sub.AddDuration(duration)
		}

		if err := srt.WriteSRTFile(subtitles, outputFile); err != nil {
			return fmt.Errorf("could not write .srt file: %s", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(offsetCmd)

	offsetCmd.Flags().StringP(
		"duration",
		"d",
		"",
		"Offset duration to add to all timestamps (e.g. -1h4m30ms)",
	)

	offsetCmd.Flags().StringP(
		"output",
		"o",
		"",
		"Write subtitles to output file",
	)
}
