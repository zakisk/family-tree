package models

// represents a person in database
type Person struct {
	ID   int
	Name string
}

// represents a role in databse
type Role struct {
	RoleID int
	Role   string
}

// represents a relationship in database
type Relationship struct {
	RelationshipID int
	Person1ID      int
	Person2ID      int
	Person1RoleID  int
	Person2RoleID  int
}
