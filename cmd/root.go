package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/fang"
	"github.com/htekgulds/terminal-rehber/pkg/tui"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdName string = "rehber"
var configFile string = "config.yaml"

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   cmdName,
	Short: "Terminal Rehber",
	Long:  "Terminalde çalışan telefon rehberi uygulaması",
	RunE: func(cmd *cobra.Command, args []string) error {
		log, err := os.Create("output.log")
		if err != nil {
			panic(err)
		}
		defer log.Close()
		slog.SetDefault(slog.New(slog.NewTextHandler(log, &slog.HandlerOptions{})))

		model := tui.NewModel()
		if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
		return nil
	},
}

func Execute() {
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file locations are: ./"+configFile+", $HOME/.config/"+cmdName+"/"+configFile+", /etc/"+cmdName+"/"+configFile)
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")

	// Bind flag to viper key
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/" + cmdName)
		viper.AddConfigPath("/etc/" + cmdName)
		viper.SetConfigName("config")
	}

	_ = godotenv.Load()
	viper.AutomaticEnv()        // read env vars that match
	viper.SetEnvPrefix(cmdName) // REHBER_ARGNAME=...

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
