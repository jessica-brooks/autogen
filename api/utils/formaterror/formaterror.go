package formaterror

import (
	"errors"
)

//FormatError returns new error..
//Can write formatting option for json here
func FormatError(err string) error {
	return errors.New("Incorrect Details")
}
