package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logger *log.Logger

	RootCmd = &cobra.Command{
		Use:   "dd",
		Short: "datadog cli",
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("%v", err.Error())
	}
}

func init() {
	logger = log.New(os.Stderr, "", log.Lshortfile)

	cobra.OnInitialize(configurator)
}

func configurator() {
	viper.SetEnvPrefix("DD")
	viper.AutomaticEnv()
}

func handleErr(r *http.Response, err error) {
	fmt.Fprintf(os.Stderr, "ERROR:\n\n%v\n\n", err)
	fmt.Fprintf(os.Stderr, "HTTP RESPONSE:\n\n%v\n", r)
	os.Exit(1)
}
