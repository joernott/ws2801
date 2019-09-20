package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var numLEDs int
var position int
var first int
var last int
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ws2801tool",
	Short: "Simple tool to manage ws2801 LED strips",
	Long:  `This tool uses the ws2801 library to manage the ws2801 LED strip.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".ws2801tool", "config file (default is $HOME/.ws2801tool.yaml)")
	rootCmd.PersistentFlags().IntVarP(&numLEDs, "number_of_leds", "n", 0, "number of LEDs")
	rootCmd.PersistentFlags().IntVarP(&position, "position", "p", 0, "Position of the pixel.")
	rootCmd.PersistentFlags().IntVarP(&first, "first", "f", 0, "Position of the first pixel when defining a range.")
	rootCmd.PersistentFlags().IntVarP(&last, "last", "l", 0, "Position of the last pixel when defining a range.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose command line output.")
}
