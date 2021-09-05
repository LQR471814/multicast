package common

const (
	missingPrivilegeMessage = "Insufficient permissions!"
	setupRequiredMessage    = "Multicasting is not setup on this machine"
	brokenInterfaceError    = "Cannot listen on this interface!"
)

type MissingPrivileges struct{}

func (err MissingPrivileges) Error() string {
	return missingPrivilegeMessage
}

func IsMissingPrivilegeError(err error) bool {
	return err.Error() == missingPrivilegeMessage
}

type SetupRequired struct{}

func (err SetupRequired) Error() string {
	return setupRequiredMessage
}

func IsSetupReqError(err error) bool {
	return err.Error() == setupRequiredMessage
}

type BrokenInterfaceError struct{}

func (err BrokenInterfaceError) Error() string {
	return brokenInterfaceError
}

func IsBrokenIntfError(err error) bool {
	return err.Error() == brokenInterfaceError
}
