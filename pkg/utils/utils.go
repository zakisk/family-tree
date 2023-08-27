package utils

import "github.com/zakisk/family-tree/pkg/errors"

// counterpart in relationship, son to daughter, son to father, wife to husband and so on.
func GetCounterPartRole(person1Role, counterpartGender string) (string, error) {
	switch person1Role {
	case "father", "mother":
		{
			if counterpartGender == "male" {
				return "son", nil
			} else {
				return "daughter", nil
			}
		}
	case "son", "daughter":
		{
			if counterpartGender == "male" {
				return "father", nil
			} else {
				return "mother", nil
			}
		}
	case "brother", "sister":
		{
			if counterpartGender == "male" {
				return "brother", nil
			} else {
				return "sister", nil
			}
		}
	case "wife":
		return "husband", nil
	case "husband":
		return "wife", nil
	default:
		return "", errors.GenderNotSpecifiedError()
	}
}