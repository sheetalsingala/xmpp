// Copyright 2016 Sam Whited.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package xmpp

import (
	"testing"

	"mellium.im/xmpp/jid"
)

func TestDialClientPanicsIfNilContext(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Dial to panic when passed a nil context.")
		}
	}()
	DialClient(nil, "tcp", jid.MustParse("feste@shakespeare.lit"))
}

// The zero value of Dialer.ConnType should be "xmpp-client"
func TestDefaultConnType(t *testing.T) {
	var d Dialer
	if ct := d.connType(); ct != "xmpp-client" {
		t.Error("Wrong zero value for conntype; expected xmpp-client but got %s", ct)
	}
}
