package shared

// Constraint: This package should only contain constant and static data

// Table type
type Table map[string][]string

// Request form
type RequestForm struct {
	Table  string
	Entry  string
	Values []string
}

// Listen port
const Port = ":1800"

// Directory environment variable key
const Direnv = "EXMROOT"
