/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dnslookup",
	Short: "Performs dns lookup and returns the IpV4 address",
	Long: `Performs dns lookup on the given dns like:
					dnslookup xyz.com
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 && len(args) <= 10 {
			return nil
		}
		return errors.New("Can accept only 1 ( min ) - 10 ( max ) domains to lookup")
	},
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			addr, err := net.LookupIP(args[i])
			if err != nil {
				fmt.Println("Cannot find IP instance for :" + args[i])
			}
			for _, ip := range addr {
				if ipv4 := ip.To4(); ipv4 != nil {
					fmt.Println(args[i] + " : " + ipv4.To4().String())
				}
			}
			fmt.Println()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dnslookup.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
