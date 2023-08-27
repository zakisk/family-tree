package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zakisk/family-tree/database"
	"github.com/zakisk/family-tree/pkg/errors"
	"github.com/zakisk/family-tree/pkg/utils"
)

var (
	flagAs        string
	flagOf        string
	person2Gender string
)

// represents `family-tree connect role [role-name] command`
var ConnectRelationshipCmd = &cobra.Command{
	Use:   "connect",
	Short: "connects two person",
	Long: `Subcommand to connect two persons with a relationship.
For example:
// connects two persons of a family
family-tree connect [person1] --as brother --of [person2]

// default gender of person2 is male if it isn't specify it with --person2-gender or -g
family-tree connect [person1] --as brother --of [person2] --person2-gender female
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

		if len(flagAs) == 0 {
			log.Fatal(errors.EmptyFlagError("as"))
			cmd.Help()
			return
		}

		if len(flagOf) == 0 {
			log.Fatal(errors.EmptyFlagError("of"))
			cmd.Help()
			return
		}

		// convert flags and args to lower case because user can pass value in any case
		person1Name := strings.ToLower(args[0])
		flagAs = strings.ToLower(flagAs)
		flagOf = strings.ToLower(flagOf)
		person2Gender = strings.ToLower(person2Gender)

		p2Role, err := utils.GetCounterPartRole(flagAs, person2Gender)
		if err != nil {
			log.Fatal(err)
			return
		}

		person1, err := db.GetPerson(person1Name)
		if err != nil {
			log.Fatal(err)
			return
		}

		person1Role, err := db.GetRole(flagAs)
		if err != nil {
			log.Fatal(err)
			return
		}

		person2, err := db.GetPerson(flagOf)
		if err != nil {
			log.Fatal(err)
			return
		}

		person2Role, err := db.GetRole(p2Role)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = db.InsertRelationship(person1.ID, person2.ID, person1Role.RoleID, person2Role.RoleID)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("%s is connected successfully to %s as %s\n", person1Name, flagOf, flagAs)
	},
}


func init() {
	ConnectRelationshipCmd.Flags().StringVarP(&flagAs, "as", "a", "", "specifies the role of the person")
	ConnectRelationshipCmd.Flags().StringVarP(&flagOf, "of", "o", "", "specifies the relation with person2")
	ConnectRelationshipCmd.Flags().StringVarP(&person2Gender, "person2-gender", "g", "male", "specifies gender of person2")
}
