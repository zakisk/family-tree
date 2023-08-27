package database

import (
	"database/sql"
	"fmt"

	"github.com/zakisk/family-tree/models"
	"github.com/zakisk/family-tree/pkg/errors"
)

// FamilyDB struct
type FamilyDB struct {
	DB *sql.DB
}

func NewFamilyDB(db *sql.DB) *FamilyDB {
	return &FamilyDB{DB: db}
}

// Get Person by name
func (db *FamilyDB) GetPerson(name string) (*models.Person, error) {
	rows, err := db.DB.Query("SELECT * FROM Persons WHERE person_name = ?", name)
	if err != nil {
		return nil, errors.GetEntityError("person", err)
	}
	defer rows.Close()

	// if person is found then it is guaranteed that there will be only one person
	person := &models.Person{}
	if rows.Next() {
		err := rows.Scan(&person.ID, &person.Name)
		if err != nil {
			return nil, errors.EntityReadError(err)
		}
	} else {
		return nil, errors.EntityNotFoundError("person", err)
	}

	return person, nil
}

func (db *FamilyDB) GetRole(roleName string) (*models.Role, error) {
	role, err := db.getRole(roleName)
	if err != nil {
		// let assume a case that we've added two persons A and B, and we're connecting them
		// `family-tree connect A --as father --of B` but as we're getting B's role in getPerson2Role function
		// it will return 'son' and we've not added role 'son' yet so it will cause an error that's why role of B
		// is inserted into databse here
		err = db.InsertRole(roleName)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Role \"%s\" is created\n", roleName)
		role, err = db.getRole(roleName)
		if err != nil {
			return nil, err
		}
	}
	return role, err
}

// Get Role by name internal
func (db *FamilyDB) getRole(roleName string) (*models.Role, error) {
	rows, err := db.DB.Query("SELECT * FROM Roles WHERE role = ?", roleName)
	if err != nil {
		return nil, errors.GetEntityError("role", err)
	}
	defer rows.Close()

	// if role is found then it is guaranteed that there will be only one role
	role := &models.Role{}
	if rows.Next() {
		err := rows.Scan(&role.RoleID, &role.Role)
		if err != nil {
			return nil, errors.EntityReadError(err)
		}
	} else {
		return nil, errors.EntityNotFoundError("role", err)
	}

	return role, nil
}

// Get relationships by person's id
func (db *FamilyDB) GetRelativesCountsByPersonID(personID, personRoleID, relativeRoleID int) (int, error) {
	query := `
		SELECT COUNT(*) FROM Relationships
		WHERE 
			(person1_id = ? AND person1_role_id = ? AND person2_role_id = ?)
			OR
			(person2_id = ? AND person2_role_id = ? AND person1_role_id = ?)
	`
	var count int
	err := db.DB.QueryRow(query, personID, personRoleID, relativeRoleID, personID, personRoleID, relativeRoleID).Scan(&count)
	if err != nil {
		return -1, errors.GetEntityError("relationships", err)
	}
	return count, nil
}

// inserts person into database, here for the sake of simplicity we're inserting only name
func (db *FamilyDB) InsertPerson(name string) error {
	_, err := db.DB.Exec("INSERT INTO Persons (person_name) VALUES (?)", name)
	if err != nil {
		return errors.InserEntityError("person", err)
	}
	return nil
}

// inserts a role into database e.g father, mother, son, daughter
func (db *FamilyDB) InsertRole(role string) error {
	_, err := db.DB.Exec("INSERT INTO Roles (role) VALUES (?)", role)
	if err != nil {
		return errors.InserEntityError("role", err)
	}
	return nil
}

// inserts a relationship between two persons and their roles ids, e.g Person_A (Father) Person_B(Son)
func (db *FamilyDB) InsertRelationship(person1_id, person2_id, person1_role_id, person2_role_id int) error {
	if !db.checkRelationshipNotExists(person1_id, person2_id, person1_role_id, person2_role_id) {
		return errors.EntityAlreadyExistsError("relationship")
	}
	_, err := db.DB.Exec(`
		INSERT INTO Relationships (person1_id, person2_id, person1_role_id, person2_role_id) VALUES (?, ?, ?, ?)`,
		person1_id, person2_id, person1_role_id, person2_role_id,
	)
	if err != nil {
		return errors.InserEntityError("relationship", err)
	}
	return nil
}

func (db *FamilyDB) checkRelationshipNotExists(person1_id, person2_id, person1_role_id, person2_role_id int) bool {
	query := `
		SELECT COUNT(*) FROM Relationships
		WHERE 
			(person1_id = ? AND person2_id = ? AND person1_role_id = ? AND person2_role_id = ?)
			OR
			(person1_id = ? AND person2_id = ? AND person1_role_id = ? AND person2_role_id = ?)
	`
	var count int
	err := db.DB.QueryRow(query, person1_id, person2_id, person1_role_id, person2_role_id,
		person2_id, person1_id, person2_role_id, person1_role_id).Scan(&count)
	if err != nil {
		return false
	}
	return count == 0
}

func (db *FamilyDB) GetRelativeByID(personID, relativeRoleID int) (*models.Person, error) {
	r := &models.Relationship{}
	query := `
		SELECT * FROM Relationships
		WHERE
			(person1_id = ? AND person2_role_id = ?)
			OR
			(person2_id = ? AND person1_role_id = ?)
	`
	row := db.DB.QueryRow(query, personID, relativeRoleID, personID, relativeRoleID)
	err := row.Scan(&r.RelationshipID, &r.Person1ID, &r.Person2ID, &r.Person1RoleID, &r.Person2RoleID)
	if err != nil {
		return nil, errors.PersonRelationshipsNotExists(err)
	}
	var relativeID int
	if r.Person1RoleID == relativeRoleID {
		relativeID = r.Person1ID
	} else {
		relativeID = r.Person2ID
	}
	person := &models.Person{}
	err = db.DB.QueryRow("SELECT * FROM Persons WHERE person_id = ?", relativeID).Scan(&person.ID, &person.Name)
	return person, nil
}