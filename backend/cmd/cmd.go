package cmd

import (
	"fmt"
	"horo/service"
	"os"

	"github.com/spf13/cobra"
)

var ConfigPath string

var RootCmd = &cobra.Command{
	Use:   "horo <command>",
	Short: "horo command for user",
}

func Execute() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "add timer",
		Run: func(cmd *cobra.Command, args []string) {
			addTimer(args)
		},
	})
	RootCmd.AddCommand(&cobra.Command{
		Use:   "del",
		Short: "delete timer",
		Run: func(cmd *cobra.Command, args []string) {
			delTimer(args)
		},
	})
	RootCmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "show timer",
		Run: func(cmd *cobra.Command, args []string) {
			showTimer()
		},
	})
	RootCmd.AddCommand(&cobra.Command{
		Use:   "daemon",
		Short: "daemon run timer",
		Run: func(cmd *cobra.Command, args []string) {
			daemon()
		},
	})

	RootCmd.PersistentFlags().StringVar(&ConfigPath, "config", ConfigPath, "set config file path.")
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addTimer(args []string) {
	service.AddTimer(args)
}

func delTimer(args []string) {
	fmt.Println("not support")
}

func showTimer() {
	service.ShowTimer()
}

func daemon() {
	//service.DaemonRun()
	fmt.Println("not support")
}
