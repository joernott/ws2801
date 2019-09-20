package cmd

import (
	"fmt"
	"os"

	"github.com/joernott/ws2801"
	"github.com/spf13/cobra"
)

var red uint8
var green uint8
var blue uint8

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets one or more pixels to a specific color",
	Long:  `ws2801tool set sets a pixel or range of pixels to a specific color`,
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
				fmt.Printf("Set single pixel %v to (r=%v, g=%v, b=%v).\n", position, red, green, blue)
			}
			err = p.SetPixel(position, red, green, blue)
		} else {
			if verbose {
				fmt.Printf("Set pixels between %v and %v to (r=%v, g=%v, b=%v).\n", first, last, red, green, blue)
			}
			err = p.SetPixels(first, last, red, green, blue)
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
	rootCmd.AddCommand(setCmd)

	setCmd.PersistentFlags().Uint8VarP(&red, "red", "r", 0, "Intensity of the red subpixel")
	setCmd.PersistentFlags().Uint8VarP(&green, "green", "g", 0, "Intensity of the green subpixel")
	setCmd.PersistentFlags().Uint8VarP(&blue, "blue", "b", 0, "Intensity of the blue subpixel")
}
