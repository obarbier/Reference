/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
  "log"
  "github.com/spf13/cobra"
  homedir "github.com/mitchellh/go-homedir"
  "database/sql"
  "github.com/spf13/viper"
  _ "github.com/mattn/go-sqlite3"

)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "todoCLI",
  Short: "Very simple Todo App to learn Cobra",
  Long: `todoCLI will allow you to create read update delete
        task using a command line build on Cobra with GOLANG.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  PreRunE:preFlight,

  RunE:func(cmd *cobra.Command, args []string) error {
    return nil
  },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
func preFlight(cmd *cobra.Command, args []string) error{

  db, err := sql.Open("sqlite3", "./database.db")

  // if there is an error opening the connection, handle it
  if err != nil {
      panic(err.Error())
  }
  // defer the close till after the main function has finished
  // executing
  defer db.Close()

    sqlStmt := `DROP table IF EXISTS TASK; -- "OR REPLACE"
    CREATE table TASK(ID INTEGER PRIMARY KEY AUTOINCREMENT, DESCRIPTION text, CREATED_DATE DATETIME DEFAULT CURRENT_TIMESTAMP, UPDATED_DATE DATETIME, STATUS text);`

    _, err = db.Exec(sqlStmt)
    if err != nil {
      // log.Printf("%q: %s\n", err, sqlStmt)
      log.Printf("%q: %s\n", err, sqlStmt)
      // return
    }
  return nil
}
func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todoCLI.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  // rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".todoCLI" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".todoCLI")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}
