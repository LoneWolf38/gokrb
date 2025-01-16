package client

import (
	"testing"

	"github.com/LoneWolf38/gokrb/config"
	"github.com/LoneWolf38/gokrb/keytab"
)

func TestAssumePreauthentication(t *testing.T) {
	t.Parallel()

	cl := NewWithKeytab("username", "REALM", &keytab.Keytab{}, &config.Config{}, AssumePreAuthentication(true))
	if !cl.settings.assumePreAuthentication {
		t.Fatal("assumePreAuthentication should be true")
	}
	if !cl.settings.AssumePreAuthentication() {
		t.Fatal("AssumePreAuthentication() should be true")
	}
}
