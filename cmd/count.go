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

var personGender string

// represents `family-tree connect role [role-name] command`
var CountRelativesCmd = &cobra.Command{
	Use:   "count",
	Short: "counts relatives",
	Long: `Subcommand to count relatives by their roles.
For example:
// count sons
family-tree count sons --of [person]

// count daughters
family-tree count daughters --of [person]
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal(errors.EmptyArgError("relative role"))
		}

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
		relatives := strings.ToLower(args[0])
		flagOf = strings.ToLower(flagOf)
		personGender = strings.ToLower(personGender)
		relativeRoleName, err := validateAndResolveArg(relatives)
		if err != nil {
			log.Fatal(err)
			return
		}

		relativeRole, err := db.GetRole(relativeRoleName)
		if err != nil {
			log.Fatal(err)
			return
		}

		// family-tree count sons --of A, A is the person name
		person, err := db.GetPerson(flagOf)
		if err != nil {
			log.Fatal(err)
			return
		}

		pRole, err := utils.GetCounterPartRole(relativeRoleName, personGender)
		if err != nil {
			log.Fatal(err)
			return
		}

		personRole, err := db.GetRole(pRole)
		if err != nil {
			log.Fatal(err)
			return
		}

		count, err := db.GetRelativesCountsByPersonID(person.ID, personRole.RoleID, relativeRole.RoleID)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("%s has %d %s\n", person.Name, count, relatives)
	},
}

func validateAndResolveArg(arg string) (string, error) {
	switch arg {
	case "sons":
		return "son", nil
	case "daughters":
		return "daughter", nil
	case "wives":
		return "wife", nil
	default:
		return "", errors.WrongRelativeArgError(arg)
	}
}

func init() {
	CountRelativesCmd.Flags().StringVarP(&flagOf, "of", "o", "", "specifies the person")
	CountRelativesCmd.Flags().StringVarP(&personGender, "person-gender", "g", "male", "specifies person gender")
}
