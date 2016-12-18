//Errcode package is a lightweight library for coded errors in Golang.
//It allows you to created and compare coded errors, which also conform
//to the regular error interface.
package errorcode

//Errorcode satisfies the stand error interface.
type Errorcode struct {
	errType    string
	errMessage string
}

var _ error = &Errorcode{}

//New creates a new coded error with a type and error message.
func New(errType, message string) *Errorcode {
	return &Errorcode{errType: errType, errMessage: message}
}

//Type returns the type of the error.
func (e *Errorcode) Type() string {
	return e.errType
}

//Error returns the error message.
func (e *Errorcode) Error() string {
	return e.errMessage
}

//SameType checks if two errors have the same type.
func (e1 *Errorcode) SameType(e2 *Errorcode) (same bool) {
	if e1.Type() == e2.Type() {
		return true
	}
	return false
}
