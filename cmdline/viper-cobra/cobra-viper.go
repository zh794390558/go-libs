package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	mainCmd.AddCommand(versionCmd)

	viper.SetEnvPrefix("DISPATH")
	viper.AutomaticEnv()

	/*
		          When AutomaticEnv called, Viper will check for an environment variable any
			        time a viper.Get request is made. It will apply the following rules. It
				      will check for a environment variable with a name matching the key
				            uppercased and prefixed with the EnvPrefix if set.
	*/

	flags := mainCmd.Flags()

	flags.Bool("debug", false, "Turen on debuggin.")
	flags.String("addr", "localhost:5002", "Address of the service")
	flags.String("smtp-addr", "localhost:25", "Address of the SMTP server")
	flags.String("smtp-user", "", "User to authenticate with the SMTP server")
	flags.String("smtp-password", "", "Password to authenticate with the SMTP server")
	flags.String("email-from", "noreply@example.com", "The from email address.")

	viper.BindPFlag("debug", flags.Lookup("debug"))
	viper.BindPFlag("addr", flags.Lookup("addr"))
	viper.BindPFlag("smtp_addr", flags.Lookup("smtp-addr"))
	viper.BindPFlag("smtp_user", flags.Lookup("smtp-user"))
	viper.BindPFlag("smtp_password", flags.Lookup("smtp-password"))
	viper.BindPFlag("email_from", flags.Lookup("email-from"))

	// Viper supports reading from yaml, toml and/or json files. Viper can
	// search multiple paths. Paths will be searched in the order they are
	// provided. Searches stopped once Config File found.

	viper.SetConfigName("conf") //name of config file(without extension)
	viper.AddConfigPath("/tmp") //path to look for the config file in
	viper.AddConfigPath(".")    // more path to look for the config files

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No config file found. Using built-in defaults.")
	}
}

// The main command describes the service and defaults to printing the
// help message.
var mainCmd = &cobra.Command{
	Use:   "dispath",
	Short: "Event dispath service.",
	Long:  `HTTP service that consumes events and dispatched them to subscribers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main")
	},
}

var version string = "1.1.0"

// The version command prints this service.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the dispatch service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		panic(err)
	}
}
