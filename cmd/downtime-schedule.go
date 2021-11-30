package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/spf13/cobra"
)

var (
	scheduleMsg   string
	scheduleTags  []string
	scheduleScope []string

	scheduleDowntimeCmd = &cobra.Command{
		Use:   "schedule",
		Short: "schedule downtime",
		Run:   scheduleDowntimeRunFunc,
		Long: `
Schedule downtime.

Usage
-----

$ dd downtime schedule -m "scheduled maintenance" -t "team:ops,..."
$ dd downtime schedule -m "scheduled maintenance" -t "team:ops,..." -s "env:prod,..."
`,
	}
)

func init() {
	downtimeCmd.AddCommand(scheduleDowntimeCmd)

	scheduleDowntimeCmd.Flags().StringVarP(&scheduleMsg, "message", "m", "", "message to include with downtime notification")
	scheduleDowntimeCmd.Flags().StringSliceVarP(&scheduleTags, "tags", "t", []string{}, "comma-separated list of monitor tags")
	scheduleDowntimeCmd.Flags().StringSliceVarP(&scheduleScope, "scope", "s", []string{}, "comma-separated list of monitor scopes")
}

func scheduleDowntimeRunFunc(cmd *cobra.Command, args []string) {
	ctx := datadog.NewDefaultContext(context.Background())
	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	if len(scheduleMsg) == 0 {
		scheduleMsg = "scheduled by datadog cli"
	}

	body := *datadog.NewDowntime()
	body.Message = &scheduleMsg
	body.MonitorTags = &scheduleTags
	body.Scope = &scheduleScope

	res, r, err := client.DowntimesApi.CreateDowntime(ctx, body)
	if err != nil {
		handleErr(r, err)
	}

	j, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("%s\n", j)
}
