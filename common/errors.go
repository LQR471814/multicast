package common

type MissingPrivileges struct{}

func (err MissingPrivileges) Error() string {
	return "Program has insufficient permissions to continue"
}

type SetupRequired struct{}

func (err SetupRequired) Error() string {
	return "Multicasting is not setup on this machine"
}
