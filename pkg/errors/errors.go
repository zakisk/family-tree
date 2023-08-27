package errors

import "fmt"

// Database errors
func InvalidDatabaseDriverError(err error) error {
	return fmt.Errorf("Failed to establish database connection \nmessage: %s", err.Error())
}

func InvalidQueryError(err error) error {
	return fmt.Errorf("Unable to query database \n%s", err.Error())
}

func TableCreationError(tableName string, err error) error {
	return fmt.Errorf("Unable to create %s table \nmessage: %s", tableName, err.Error())
}

func InserEntityError(entityName string, err error) error {
	return fmt.Errorf("Unable to insert %s\nmessage: %s", entityName, err)
}

func GetEntityError(entityName string, err error) error {
	return fmt.Errorf("Unable to get %s\nmessage: %s", entityName, err)
}

func EntityReadError(err error) error {
	return fmt.Errorf("Unable to read fields from result\nmessage: %s", err)
}

func EntityNotFoundError(role string ,err error) error {
	return fmt.Errorf("no %s is found\nmessage: %s", role, err)
}

func EntityAlreadyExistsError(entity string) error {
	return fmt.Errorf("%s already exists\n", entity)
}

func PersonRelationshipsNotExists(err error) error {
	return fmt.Errorf("Person has no such relationship\nmessage: %s", err)
}


// Command errors
func EmptyArgError(argName string) error {
	return fmt.Errorf("Please spcify %s", argName)
}

func EmptyFlagError(flagName string) error {
	return fmt.Errorf("Please spcify flag \"%s\"", flagName)
}

func GenderNotSpecifiedError() error {
	return fmt.Errorf("Please spcify correct gender")
}

func WrongRelativeArgError(arg string) error {
	return fmt.Errorf("%s is not a role. allowed values [sons, daughters, wives]", arg)
}