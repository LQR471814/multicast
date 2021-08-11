package common

type MissingPrivileges struct{}

func (err MissingPrivileges) Error() string {
	return "Program has insufficient permissions to continue"
}
