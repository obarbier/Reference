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
	"strings"
	_"time"
	"log"
	"database/sql"
	"github.com/spf13/cobra"
)
type Task struct {
    Description string `json:"description"`
		status string `json:"status"`
}
var (

// createTaskCmd represents the createTask command
 createTaskCmd = &cobra.Command{
	Use:   "createTask",
	Short: "Log task into an SQLITE table ",
	Long: ` For example:
	todoCLI createTask read 5 pages of google SRE book.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "./database.db")

		// if there is an error opening the connection, handle it
		if err != nil {
				panic(err.Error())
		}
		// defer the close till after the main function has finished
		// executing
		defer db.Close()
			task := Task{  strings.Join(args, " "),  "ACTIVE"}
			sqlStmt:=fmt.Sprintf("INSERT INTO TASK( DESCRIPTION  ,UPDATED_DATE, STATUS )VALUES  ('%s' ,CURRENT_TIMESTAMP, '%s');",task.Description,  task.status  )
			_, err = db.Exec(sqlStmt)
			if err != nil {
				// log.Printf("%q: %s\n", err, sqlStmt)
				log.Printf("%q: %s\n", err, sqlStmt)
				// return
			}
		return

	},
}

)

func init() {
	rootCmd.AddCommand(createTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// createTaskCmd.Flags().StringVarP(&task, "task", "t", "", "what is your task")
}
