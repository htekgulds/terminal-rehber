package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var name string

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello to someone",
	Run: func(cmd *cobra.Command, args []string) {
		verbose := viper.GetBool("verbose")
		if verbose {
			fmt.Println("Verbose mode is ON")
		}
		name := viper.GetString("name")
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	helloCmd.Flags().StringVarP(&name, "name", "n", "world", "name to greet")
	viper.BindPFlag("name", helloCmd.Flags().Lookup("name"))
}
