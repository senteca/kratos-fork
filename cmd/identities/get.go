package identities

import (
	"fmt"
	"time"

	"github.com/ory/kratos/internal/clihelpers"
	"github.com/ory/kratos/internal/httpclient/models"

	"github.com/spf13/cobra"

	"github.com/ory/kratos/cmd/cliclient"
	"github.com/ory/kratos/internal/httpclient/client/admin"
	"github.com/ory/x/cmdx"
)

var getCmd = &cobra.Command{
	Use:   "get <id-0 [id-1 ...]>",
	Short: "Get one or more identities by ID",
	Long: fmt.Sprintf(`This command gets all the details about an identity. To get an identity by some selector, e.g. the recovery email address, use the list command in combination with jq.
Example: get the identities with the recovery email address at the domain "ory.sh":

kratos identities get $(kratos identities list --format json | jq -r 'map(select(.recovery_addresses[].value | endswith("@ory.sh"))) | .[].id')

%s
`, clihelpers.WarningJQIsComplicated),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := cliclient.NewClient(cmd)

		identities := make([]*models.Identity, 0, len(args))
		for _, id := range args {
			resp, err := c.Admin.GetIdentity(admin.NewGetIdentityParamsWithTimeout(time.Second).WithID(id))
			cmdx.Must(err, "Could not get identity \"%s\": %s", args[0], err)

			identities = append(identities, resp.Payload)
		}

		if len(identities) == 1 {
			clihelpers.PrintRow(cmd, (*outputIdentity)(identities[0]))
			return
		}
		clihelpers.PrintCollection(cmd, &outputIdentityCollection{identities})
	},
}
