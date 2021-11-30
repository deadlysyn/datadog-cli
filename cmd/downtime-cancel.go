package cmd

import (
	"context"
	"fmt"
	"sort"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/spf13/cobra"
)

var (
	cancelID     int64
	cancelScopes string
	cancelTags   []string

	cancelDowntimeCmd = &cobra.Command{
		Use:   "cancel",
		Short: "cancel downtime",
		Run:   cancelDowntimeRunFunc,
		Long: `
Cancel downtime.

Usage
-----

$ dd downtime cancel -i 1516340265
$ dd downtime cancel -s env:prod
$ dd downtime cancel -t team:ops,service:foo
`,
	}
)

func init() {
	downtimeCmd.AddCommand(cancelDowntimeCmd)

	cancelDowntimeCmd.Flags().Int64VarP(&cancelID, "id", "i", 0, "downtime ID")
	cancelDowntimeCmd.Flags().StringVarP(&cancelScopes, "scopes", "s", "", "comma-delimited list of downtime scopes")
	cancelDowntimeCmd.Flags().StringSliceVarP(&cancelTags, "tags", "t", []string{}, "comma-delimited list of monitor tags")
}

func cancelDowntimeRunFunc(cmd *cobra.Command, args []string) {
	if cancelID > 0 {
		cancelByID(cancelID)
	} else if len(cancelScopes) != 0 {
		cancelByScope(cancelScopes)
	} else if len(cancelTags) != 0 {
		cancelByTag(cancelTags)
	}
}

func cancelByID(ID int64) {
	ctx := datadog.NewDefaultContext(context.Background())
	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	r, err := client.DowntimesApi.CancelDowntime(ctx, ID)
	if err != nil {
		handleErr(r, err)
	}
	fmt.Printf("cancelled downtime %v\n", ID)
}

func cancelByScope(scopes string) {
	ctx := datadog.NewDefaultContext(context.Background())
	cfg := datadog.NewConfiguration()
	client := datadog.NewAPIClient(cfg)

	body := *datadog.NewCancelDowntimesByScopeRequest(scopes)
	res, r, err := client.DowntimesApi.CancelDowntimesByScope(ctx, body)
	if err != nil {
		handleErr(r, err)
	}
	fmt.Printf("cancelled downtime(s) %v\n", res.CancelledIds)
}

func cancelByTag(tags []string) {
	res := getDowntimes()

	var ids []int64
	for _, d := range res {
		if slicesEqual(tags, *d.MonitorTags) {
			ids = append(ids, *d.Id)
		}
	}

	if len(ids) == 0 {
		fmt.Println("no matching downtimes found")
	} else {
		for _, v := range ids {
			cancelByID(v)
		}
	}
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.SliceStable(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.SliceStable(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
