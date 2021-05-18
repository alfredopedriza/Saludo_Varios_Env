/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// despedidaCmd represents the despedida command
var despedidaCmd = &cobra.Command{
	Use:   "despedida",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		//viper.AutomaticEnv()
		//viper.SetEnvPrefix("PREFIX")
		//viper.BindEnv("name", "buenosdias")

		viper.SetConfigName(".adios") // name of config file (without extension)
		viper.AddConfigPath(".")      // optionally look for config in the working directory
		err := viper.ReadInConfig()   // Find and read the config file
		if err != nil {               // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}

		//viper.SetEnvPrefix("PRUEBA2")
		//viper.BindEnv("name", "buenosdias")
		//fmt.Println("Hello World")
		/*var name string
		var verbose bool
		name, _ = rootCmd.Flags().GetString("name")
		verbose, _ = rootCmd.Flags().GetBool("verbose")
		*/
		if !verbose {
			verbose = viper.GetBool("verbose")
		}

		if name == "" {

			name = viper.GetString("name")
		}

		fmt.Println("Adios", name)
		fmt.Println("verbose", verbose)

	},
}

func init() {
	rootCmd.AddCommand(despedidaCmd)

	//viper.SetEnvPrefix("PREFIX")
	despedidaCmd.Flags().StringVar(&name, "name", "", "name for copyright attribution")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// despedidaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// despedidaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
