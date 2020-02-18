package errs

import "errors"

var (
    InvalidVolume = "Invalid Volume"
    InvalidNode   = "Invalid Node"
    AlreadyExists = "Already Exists"
)

func New(err string) (error) {
    return errors.New(err)
}
