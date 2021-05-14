package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	thegraph "github.com/streamingfast/thegraphhose"
	"gopkg.in/yaml.v3"
)

var startCmd = &cobra.Command{
	Use:   "start <subgraph-manifest> <subgraph-wasm>",
	Short: "Consume blocks from the network and feed a sub-graph aware WASM file, outputs JSON from handled event",
	Args:  cobra.ExactArgs(2),
	RunE:  startE,
}

func init() {
	rootCmd.AddCommand(startCmd)

	// serveCmd.Flags().StringP("secret", "s", "", "The accepted Bearer secret value, used to restrict access to known component")
	// serveCmd.Flags().StringP("listen-addr", "l", ":8080", "HTTP address to listen for requests")
}

func startE(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	manifestFilePath := args[0]
	wasmFilePath := args[1]

	manifestFile, err := os.Open(manifestFilePath)
	if err != nil {
		return fmt.Errorf("unable to read manifest: %w", err)
	}

	manifest := &thegraph.Manifest{}
	if err := yaml.NewDecoder(manifestFile).Decode(manifest); err != nil {
		return fmt.Errorf("unable to decode manifest: %w", err)
	}

	env := &thegraph.DefaultEnvironment{}

	_, err = thegraph.Simulate(env, wasmFilePath, "handleNewPair", nil)
	return err
}
