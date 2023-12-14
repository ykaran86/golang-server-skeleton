package cmd

import (
	"github.com/spf13/cobra"
	"golang-server-skeleton/server/api"
)

func NewCLICommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "Blog App",
	}
	cmd.AddCommand(newStartServerCommand())
	return cmd
}

func newStartServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start Blog App API Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			api.NewServer().Start()
			return nil
		},
	}
}
