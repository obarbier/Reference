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
	"strings"
	"bytes"
	"database/sql"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var (
	update_id int;
	update_switch bool;
	updateCmd = &cobra.Command{
	Use:   "update",
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
		var sqlStmt bytes.Buffer
		sqlStmt.WriteString( "UPDATE TASK SET ")
		update := make([]string, 0)
		if update_switch {
			update = append(update, fmt.Sprintf(`status =  CASE
											 			WHEN status = 'ACTIVE' THEN 'DONE'
											 			WHEN status = 'DONE' THEN 'ACTIVE'
											 			ELSE status
									 					END `))
		}

		if strings.Join(args," ") != ""{
			update = append(update,fmt.Sprintf( "description = '%s' ",strings.Join(args," ") ))
		}
		if strings.Join(update,",") == "" {
			fmt.Printf("Nothing to update")
			return
		}else{
			sqlStmt.WriteString(strings.Join(update,","))
			sqlStmt.WriteString(fmt.Sprintf(	", UPDATED_DATE = CURRENT_TIMESTAMP  WHERE id = %d;", update_id))
			_, err = db.Exec(sqlStmt.String())
			if err != nil {
				// log.Printf("%q: %s\n", err, sqlStmt)
				log.Printf("%q: %s\n", err, sqlStmt)
				// return
			}
			fmt.Printf("Task ID: %d was updated",update_id)
			return
		}

	},

})

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	 updateCmd.Flags().IntVarP(&update_id, "id","i", 0, "id to return, by default will return all task")
	 updateCmd.MarkFlagRequired("id")
	 updateCmd.Flags().BoolVarP(&update_switch,"switch", "s", false, "Help message for toggle")
}
