package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zakisk/family-tree/pkg/errors"
)

func GetDatabase() (*FamilyDB, error) {
	db, err := sql.Open("sqlite3", "./database/family-database.db")
	if err != nil {
		return nil, errors.InvalidDatabaseDriverError(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Persons (
			person_id INTEGER PRIMARY KEY NOT NULL,
			person_name TEXT UNIQUE
		)
	`)

	if err != nil {
		return nil, errors.TableCreationError("Persons", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Roles (
			role_id INTEGER PRIMARY KEY NOT NULL,
			role TEXT UNIQUE,
			CHECK (role IN ('father', 'mother', 'son', 'daughter', 'brother', 'sister', 'wife', 'husband'))
		)
	`)

	if err != nil {
		return nil, errors.TableCreationError("Roles", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Relationships (
			relationship_id INTEGER PRIMARY KEY NOT NULL,
			person1_id INTEGER,
			person2_id INTEGER,
			person1_role_id INTEGER,
			person2_role_id INTEGER,
			FOREIGN KEY (person1_id) REFERENCES Persons(person_id),
			FOREIGN KEY (person2_id) REFERENCES Persons(person_id),
			FOREIGN KEY (person1_role_id) REFERENCES Roles(role_id),
			FOREIGN KEY (person2_role_id) REFERENCES Roles(role_id),
			CONSTRAINT unique_relationship UNIQUE(person1_id, person2_id, person1_role_id, person2_role_id),
			CHECK (person1_id <> person2_id),
			CHECK (person1_role_id <> person2_role_id)
		)
	`)

	if err != nil {
		return nil, errors.TableCreationError("Relationships", err)
	}

	familyDb := NewFamilyDB(db)

	return familyDb, nil
}