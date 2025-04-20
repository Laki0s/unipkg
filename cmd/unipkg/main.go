package main

import (
	"fmt"
	"unipkg/drivers"
	"unipkg/internal/logger"
	"unipkg/pkg/distro"

	"github.com/spf13/cobra"
)

func main() {
	logger.SetVerbose(false)
	root := &cobra.Command{
		Use:   "unipkg",
		Short: "Unified package manager",
	}
	root.AddCommand(installCmd(), removeCmd())
	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}

func installCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install [pkg]",
		Short: "Install a package",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			drv, err := selectDriver()
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := drv.Install(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func removeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove [pkg]",
		Short: "Remove a package",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			drv, err := selectDriver()
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := drv.Remove(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}
}

func selectDriver() (drivers.Driver, error) {
	d := distro.Detect()
	logger.Debug("Detected distro: %s", d)
	drv, err := drivers.Select(d)
	if err == nil {
		return drv, nil
	}
	return drivers.Fallback()
}