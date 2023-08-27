package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zakisk/family-tree/database"
	"github.com/zakisk/family-tree/pkg/errors"
)

// represents `family-tree connect role [role-name] command`
var FatherCmd = &cobra.Command{
	Use:   "father",
	Short: "shows father name",
	Long: `Subcommand to find father name.
For example:
family-tree father --of [person]
`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.GetDatabase()
		defer db.DB.Close()
		if err != nil {
			log.Fatal(err)
			return
		}

		if len(flagOf) == 0 {
			log.Fatal(errors.EmptyFlagError("of"))
			cmd.Help()
			return
		}

		// convert flags and args to lower case because user can pass value in any case
		flagOf = strings.ToLower(flagOf)
		
		person, err := db.GetPerson(flagOf)
		if err != nil {
			log.Fatal(err)
			return
		}

		fatherRole, err := db.GetRole("father")
		if err != nil {
			log.Fatal(err)
			return
		}

		father, err := db.GetRelativeByID(person.ID, fatherRole.RoleID)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("Father of %s is %s\n", flagOf, father.Name)
	},
}

func init() {
	FatherCmd.Flags().StringVarP(&flagOf, "of", "o", "", "specifies person name")
}
