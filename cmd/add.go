package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zakisk/family-tree/database"
	"github.com/zakisk/family-tree/pkg/errors"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "adds names and roles",
	Long: `Subcommand to add names of family members and roles of them. 
For example:
// adds a person in family
family-tree add person john

// adds an role entry e.g father, mother, son, daughter
family-tree add role father
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			cmd.Help()
		}
		return nil
	},
}

// represents `family-tree add person [person-name] command`
var addPersonCmd = &cobra.Command{
	Use:   "person",
	Short: "adds person",
	Long: `Subcommand to add persons (duplicate not allowed). 
For example:
// adds a person in family
family-tree add person [person-name]
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal(errors.EmptyArgError("person name"))
		}

		db, err := database.GetDatabase()
		defer db.DB.Close()
		if err != nil {
			log.Fatal(err)
			return
		}

		personName := strings.ToLower(args[0])
		err = db.InsertPerson(personName)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("%s is added successfully\n", personName)
	},
}

// represents `family-tree add role [role-name] command`
var addRoleCmd = &cobra.Command{
	Use:   "role",
	Short: "adds role",
	Long: `Subcommand to add roles (duplicate not allowed). 
For example:
// adds role of a person
// allowed roles [father, mother, son, daughter, brother, sister, wife, husband]
family-tree add role [role-name]
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal(errors.EmptyArgError("role"))
		}

		// convert role to lower case because user can pass any value e.g Father or FATHER or fAtHeR
		roleName := strings.ToLower(args[0])
		db, err := database.GetDatabase()
		defer db.DB.Close()
		if err != nil {
			log.Fatal(err)
			return
		}

		err = db.InsertRole(roleName)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("%s is added successfully\n", roleName)
	},
}


func init() {
	AddCmd.AddCommand(addPersonCmd, addRoleCmd)
}
