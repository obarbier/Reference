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
	"log"
	"database/sql"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var(
delete_id int;
deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./database.db")

		// if there is an error opening the connection, handle it
		if err != nil {
				panic(err.Error())
		}
		// defer the close till after the main function has finished
		// executing
		defer db.Close()
			var sqlStmt string
			if delete_id == 0 {
				sqlStmt = "DELETE FROM TASK"
			}else{
				 sqlStmt = fmt.Sprintf("DELETE FROM TASK where id = %d",delete_id)
			}
		_, err = db.Exec(sqlStmt)
		if err != nil {
			// log.Printf("%q: %s\n", err, sqlStmt)
			log.Printf("%q: %s\n", err, sqlStmt)
			// return
		}
		fmt.Printf("Task ID: %d was delete",delete_id)
		return
	},
})

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.Flags().IntVarP(&delete_id, "id", "i", 0, "Help message for toggle")
}
