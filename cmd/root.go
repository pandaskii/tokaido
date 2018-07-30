package cmd

import (
	"bitbucket.org/ironstar/tokaido-cli/conf"
	"bitbucket.org/ironstar/tokaido-cli/system/version"

	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd - `tok`
var rootCmd = cobra.Command{
	Use:   "tok",
	Short: "Use Tokaido to bootstrap your Drupal applications",
	Long:  "Easily build out your Drupal application. Built in Go by the team at Ironstar.",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(ConfigGetCmd)
	rootCmd.AddCommand(ConfigSetCmd)
	rootCmd.AddCommand(DestroyCmd)
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(IronstarCmd)
	rootCmd.AddCommand(OpenCmd)
	rootCmd.AddCommand(PortsCmd)
	rootCmd.AddCommand(PsCmd)
	rootCmd.AddCommand(RepairCmd)
	rootCmd.AddCommand(StopCmd)
	rootCmd.AddCommand(SSHCmd)
	rootCmd.AddCommand(SyscheckCmd)
	rootCmd.AddCommand(StatusCmd)
	rootCmd.AddCommand(SyncCmd)
	rootCmd.AddCommand(UpCmd)
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(WatchCmd)
}

// Execute - Root executable
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// RootCmd will setup and return the root command
func RootCmd() *cobra.Command {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Specify the Tokaido config file to use")
	rootCmd.PersistentFlags().BoolP("force", "", false, "Forcefully skip confirmation prompts with 'yes' response")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Check the current Tokaido version (base command only)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode, command output is printed to the console")
	rootCmd.PersistentFlags().BoolP("customCompose", "", false, "When this is enabled, Tokaido will not modify the docker-compose.tok.yml file if it exists")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	conf.LoadConfig(cmd)

	if viper.GetBool("version") == true {
		fmt.Printf("v%s\n", version.Get().Version)
	} else {
		fmt.Printf("Tokaido v%s\n\n", version.Get().Version)
		fmt.Println("For help with Tokaido run `tok --help` or take a look at our documentation at https://docs.tokaido.io/")
	}
}
