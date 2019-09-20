package cmd

import (
	"fmt"
	"os"

	"github.com/joernott/ws2801"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears a pixel or a range of pixels",
	Long:  `ws2801tool clear clears a pixel or range of pixels`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootCmd.PersistentFlags().Changed("position") &&
			(rootCmd.PersistentFlags().Changed("first") ||
				rootCmd.PersistentFlags().Changed("last")) {
			fmt.Println("Either specify position or start and end.")
			os.Exit(1)
		}
		if verbose {
			fmt.Printf("Initialize strip with %v LEDs.\n", numLEDs)
		}
		p, err := ws2801.NewPixels(numLEDs)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if rootCmd.PersistentFlags().Changed("position") {
			if verbose {
				fmt.Printf("Clear single pixel %v.\n", position)
			}
			err = p.ClearPixel(position)
		} else {
			if verbose {
				fmt.Printf("Clear pixels between %v and %v.\n", first, last)
			}
			err = p.ClearPixels(first, last)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		err = p.Show()
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
