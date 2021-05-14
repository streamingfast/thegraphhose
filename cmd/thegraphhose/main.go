package main

import (
	"github.com/dfuse-io/logging"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "thegraphhose", Short: "Registrator main entry point"}

var metricsListenAddr = "localhost:9102"
var pprofListenAddr = "localhost:6060"

var zlog = logging.NewSimpleLogger("thegraphhose", "github.com/dfuse-io/thegraphhose/cmd/thegraphhose")

func init() {
	logging.Override(zlog)
}

func main() {
	setup()

	cobra.OnInitialize(func() {
		autoBind(rootCmd, "THEGRAPHHOSE")
	})

	// rootCmd.PersistentFlags().StringP("admin-dsn", "a", "inmemory://", "Admin API to use to register user, inmemory:// or {endpoint} to Admin API")
	// rootCmd.PersistentFlags().StringP("notifier-dsn", "n", "log://", "Notifier API to use to register user, log:// or mailchimp://{id},{secret}")

	rootCmd.Execute()
}
