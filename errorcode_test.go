package errorcode_test

import (
	. "github.com/EastWest14/errorcode"
	"testing"
)

func TestErrorCodes(t *testing.T) {
	const (
		type1    = "Type1"
		message1 = "Message"

		type2    = ""
		message2 = "Message1"
	)

	errorCode1 := New(type1, message1)
	if errorCode1 == nil {
		t.Error("Error code I initialized to nil, expected non-nil")
	}
	if typeReturned := errorCode1.Type(); typeReturned != type1 {
		t.Errorf("Expected type on error code I to be %s, got %s", type1, typeReturned)
	}
	if messageReturned := errorCode1.Error(); messageReturned != message1 {
		t.Errorf("Expected message on error code I to be %s, got %s", message1, messageReturned)
	}

	errorCode2 := New(type2, message2)
	if errorCode2 == nil {
		t.Error("Error code II initialized to nil, expected non-nil")
	}
	if typeReturned := errorCode2.Type(); typeReturned != type2 {
		t.Errorf("Expected type on error code II to be %s, got %s", type2, typeReturned)
	}
	if messageReturned := errorCode2.Error(); messageReturned != message2 {
		t.Errorf("Expected message on error code II to be %s, got %s", message2, messageReturned)
	}

	errorCode3 := New(type1, "Some other message, but same type")
	if !errorCode3.SameType(errorCode1) {
		t.Error("Expected error codes I and III to have the same type, got false")
	}

	if errorCode1.SameType(errorCode2) {
		t.Error("Expected error codes I and II to be of different types, got the same type")
	}
}

func Example() {
	const (
		//Error types
		OBJECT_DOESNT_EXIST_ERROR = "OBJECT_DOESNT_EXIST"
		SOME_OTHER_ERROR          = "..."
		//...
	)

	_ = func() *Errorcode {
		return New(OBJECT_DOESNT_EXIST_ERROR, "Object named '...' not found")
	}
}
