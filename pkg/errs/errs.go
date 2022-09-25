package errs

import "errors"

var (
	InvalidVolume = "Invalid Volume"
	InvalidNode   = "Invalid Node"
	AlreadyExists = "Already Exists"
	VDINotFound   = "VDI Not found"
)

func New(err string) error {
	return errors.New(err)
}
