package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

func rootCommand() *cobra.Command {
	c := &cobra.Command{
		Short: "Compose Provider Example",
		Use:   "demo",
	}
	c.AddCommand(newComposeCommand())
	return c
}

func newComposeCommand() *cobra.Command {
	c := &cobra.Command{
		Use:              "compose EVENT",
		TraverseChildren: true,
	}
	c.PersistentFlags().String("project-name", "", "compose project name") // unused
	c.AddCommand(&cobra.Command{
		Use:  "up",
		RunE: up,
		Args: cobra.ExactArgs(1),
	})
	c.AddCommand(&cobra.Command{
		Use:  "down",
		Args: cobra.ExactArgs(1),
	})
	return c
}

func main() {
	root := rootCommand()
	if plugin.RunningStandalone() {
		root.Execute()
	} else {
		cli, _ := command.NewDockerCli()

		err := plugin.RunPlugin(cli, root, manager.Metadata{
			SchemaVersion: "0.1.0",
			Vendor:        "Docker Inc.",
			Version:       "0",
		})
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func up(cmd *cobra.Command, args []string) error {
	// servicename := args[0]
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		err := infof("Running for %d\"", i)
		if err != nil {
			return err
		}
	}
	errorf("took too long to run")
	return nil
}

type jsonMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func infof(s string, args ...any) error {
	marshal, err := json.Marshal(jsonMessage{
		Type:    "info",
		Message: fmt.Sprintf(s, args...),
	})
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(marshal))
	return err
}

func errorf(s string, args ...any) error {
	marshal, err := json.Marshal(jsonMessage{
		Type:    "error",
		Message: fmt.Sprintf(s, args...),
	})
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(marshal))
	return err
}
