package cmd

import (
	"fmt"
	"github.com/lithictech/webhookdb-cli/config"
	"github.com/urfave/cli/v2"
	"os"
)

var versionCmd = &cli.Command{
	Name: "version",
	Action: func(c *cli.Context) error {
		shaPart := config.BuildSha
		if len(shaPart) >= 8 {
			shaPart = fmt.Sprintf(" (%s)", config.BuildSha[0:8])
		}
		fmt.Fprintf(os.Stdout, "%s%s\n", config.Version, shaPart)
		return nil
	},
}