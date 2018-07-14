package cmd

import (
	"bitbucket.org/ironstar/tokaido-cli/conf"
	"bitbucket.org/ironstar/tokaido-cli/services/docker"
	"bitbucket.org/ironstar/tokaido-cli/utils"

	// "fmt"

	"github.com/spf13/cobra"
)

// PsCmd - `tok ps`
var PsCmd = &cobra.Command{
	Use:   "ps",
	Short: "docker ps for your containers",
	Long:  "Alias for running `docker-compose -f <composefile> ps`",
	Run: func(cmd *cobra.Command, args []string) {
		utils.CheckCmdHard("docker-compose")
		conf.LoadConfig(cmd)

		docker.HardCheckTokCompose()

		docker.Ps()
	},
}
