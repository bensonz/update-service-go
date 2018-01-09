package communication_test

import (
	"testing"
	"../communication"
)

func singleTestDeciphor(s string, expected string, t *testing.T) {
	received := communication.Deciphor(s)
	if  received != expected {
		t.Error("Test",s,"failed")
		t.Error("Received:",received)
		t.Error("Expected:",expected)
	}
}

func TestDeciphor(t *testing.T) {
	s := ""
	singleTestDeciphor(s, "", t)
	s = "[[::a::]]"
	singleTestDeciphor(s, "a", t)
	s = "adf"
	singleTestDeciphor(s, "", t)
	s = "[[::]]"
	singleTestDeciphor(s, "", t)
	s = "[[::::]]"
	singleTestDeciphor(s, "", t)
	s = "[[::a:b:c::]]"
	singleTestDeciphor(s, "a:b:c", t)
	s = "[[::a-b/`12`[[]]1234c::]]"
	singleTestDeciphor(s, "a-b/`12`[[]]1234c", t)
        // not sure if this case is possible
	// s = "[[::a::]][[::b"
	// singleTestDeciphor(s, "a", t)

}
