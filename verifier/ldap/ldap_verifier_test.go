package ldap

import (
	"testing"

	"os"

	. "github.com/bborbe/assert"
	"github.com/bborbe/auth_http_proxy/verifier"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestImplementsVerifier(t *testing.T) {
	object := New("", "", "", 0, false, true, "", "", "", "", "", "", "", "")
	var expected *verifier.Verifier
	err := AssertThat(object, Implements(expected))
	if err != nil {
		t.Fatal(err)
	}
}

func TestServername(t *testing.T) {
	object := New("", "", "", 0, false, true, "", "", "", "", "", "", "", "")
	var expected *verifier.Verifier
	err := AssertThat(object, Implements(expected))
	if err != nil {
		t.Fatal(err)
	}
}
