package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/spf13/cobra"
)

var (
	currentOnly bool

	listDowntimeCmd = &cobra.Command{
		Use:   "list",
		Short: "list downtime",
		Run:   listDowntimeRunFunc,
		Long: `
List downtime.

Usage
-----

$ dd downtime list
$ dd downtime list --current
`,
	}
)

func init() {
	downtimeCmd.AddCommand(listDowntimeCmd)

	listDowntimeCmd.Flags().BoolVarP(&currentOnly, "current", "c", false, "only list active downtime")
}

func listDowntimeRunFunc(cmd *cobra.Command, args []string) {
	res := getDowntimes()
	j, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("%s\n", j)
}

func getDowntimes() []datadog.Downtime {
	ctx := datadog.NewDefaultContext(context.Background())

	optionalParams := datadog.ListDowntimesOptionalParameters{
		CurrentOnly: &currentOnly,
	}

	configuration := datadog.NewConfiguration()

	apiClient := datadog.NewAPIClient(configuration)
	res, r, err := apiClient.DowntimesApi.ListDowntimes(ctx, optionalParams)
	if err != nil {
		handleErr(r, err)
	}

	return res
}
