package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var key string
var secret string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Supply your Nexmo account credentials to this tool",
	Long:  `Give this tool access to your Nexmo account. To find your credentials (and sign up for an account if you don't already have one), go here: https://dashboard.nexmo.com.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		viper.Set("api-key", key)
		viper.Set("api-secret", secret)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringVarP(&key, "api-key", "k", "", "The API key of your Nexmo account")
	loginCmd.MarkFlagRequired("api-key")
	loginCmd.Flags().StringVarP(&secret, "api-secret", "s", "", "The API secret of your Nexmo account")
	loginCmd.MarkFlagRequired("api-secret")
}
