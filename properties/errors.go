package properties

import (
	_ "errors"
	"fmt"
)

type MissingPropertyError struct {
	property string
}

func (e *MissingPropertyError) Error() string {
	return fmt.Sprintf("Missing '%s' property", e.property)
}

func (e *MissingPropertyError) String() string {
	return e.Error()
}

func MissingProperty(prop string) error {
	return &MissingPropertyError{property: prop}
}

type SetPropertyError struct {
	property string
	error    error
}

func (e *SetPropertyError) Error() string {
	return fmt.Sprintf("Failed to set '%s' property, %v", e.property, e.error)
}

func (e *SetPropertyError) String() string {
	return e.Error()
}

func SetPropertyFailed(prop string, err error) error {
	return &SetPropertyError{property: prop, error: err}
}

type RemovePropertyError struct {
	property string
	error    error
}

func (e *RemovePropertyError) Error() string {
	return fmt.Sprintf("Failed to remove '%s' property, %v", e.property, e.error)
}

func (e *RemovePropertyError) String() string {
	return e.Error()
}

func RemovePropertyFailed(prop string, err error) error {
	return &RemovePropertyError{property: prop, error: err}
}
