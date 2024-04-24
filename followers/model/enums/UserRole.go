package enums

import (
	"fmt"
)

type UserRole int

const (
	Administrator UserRole = iota
	Author
	Tourist
)

func (e UserRole) String() string {
	rolesStrings := [...]string{"Administrator", "Author", "Tourist"}
	if e < Administrator || e > Tourist {
		return "Unknown"
	}
	return rolesStrings[e]
}

func (e *UserRole) FromString(s string) error {
	switch s {
	case "Administrator":
		*e = Administrator
	case "Author":
		*e = Author
	case "Tourist":
		*e = Tourist
	default:
		return fmt.Errorf("invalid UserRole: %s", s)
	}
	return nil
}
