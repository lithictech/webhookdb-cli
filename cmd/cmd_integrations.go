package cmd

import (
	"context"
	"fmt"
	"github.com/lithictech/webhookdb-cli/appcontext"
	"github.com/lithictech/webhookdb-cli/client"
	"github.com/lithictech/webhookdb-cli/prefs"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var integrationsCmd = &cli.Command{
	Name:        "integrations",
	Description: "Make sure that you're working on the correct organization when you create an integration.",
	Subcommands: []*cli.Command{
		{
			Name:        "create",
			Description: "create an integration for the given organization",
			Flags:       []cli.Flag{orgFlag()},
			Action: cliAction(func(c *cli.Context, ac appcontext.AppContext, ctx context.Context, p prefs.Prefs) error {
				if c.NArg() != 1 {
					return errors.New("Service name required. Use 'webhookdb services list' to view all available services.")
				}
				input := client.IntegrationsCreateInput{
					AuthCookie:    p.AuthCookie,
					OrgIdentifier: getOrgFlag(c, p),
					ServiceName:   c.Args().Get(0),
				}
				step, err := client.IntegrationsCreate(ctx, input)
				if err != nil {
					return err
				}
				if err := client.NewStateMachine().Run(ctx, p, step); err != nil {
					return err
				}
				return nil
			}),
		},
		{
			Name:        "list",
			Description: "list all integrations for the given organization",
			Flags:       []cli.Flag{orgFlag()},
			Action: cliAction(func(c *cli.Context, ac appcontext.AppContext, ctx context.Context, p prefs.Prefs) error {
				input := client.IntegrationsListInput{
					AuthCookie:    p.AuthCookie,
					OrgIdentifier: getOrgFlag(c, p),
				}
				out, err := client.IntegrationsList(ctx, input)
				if err != nil {
					return err
				}
				if len(out.Data) == 0 {
					// TODO: fix the bug where the cli library's timestamp shows up if we return an `errors.New()`
					return errors.New("This organization doesn't have any integrations set up yet.")
				}
				// TODO: Get this spacing correct
				fmt.Println("service name \t\t\t\t table name \t\t\t\t id")
				for _, value := range out.Data {
					fmt.Println(value.ServiceName + " \t\t\t " + value.TableName + " \t\t\t " + value.OpaqueId)
				}
				return nil
			}),
		},
	},
}
