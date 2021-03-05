package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrorCenter(t *testing.T) {
	ec := New()
	ec.Push(nil)
	ec.Push(fmt.Errorf("error1"))
	ec.Push(fmt.Errorf("error2"))

	if ec.IsEmpty() {
		t.Errorf("ec did not empty, should one error")
	}

	if strings.TrimSpace(ec.Error()) != `error1
error2` {
		t.Errorf("ec.Error() did not match, wont error1\nerror2, got %s", ec.Error())
	}
}
