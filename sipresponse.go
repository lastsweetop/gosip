package sip

import (
	"bytes"
)

type SipResponse struct {
	StateLine *StateLine
	SipHeader
	CommonHeader
	WWWAuthenticate *WWWAuthenticate
	Date            string
}

type WWWAuthenticate struct {
	Realm string
	Nonce string
	Uri   string
}

func (wa WWWAuthenticate) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("WWW-Authenticate: Digest realm=\"")
	b.WriteString(wa.Realm)
	b.WriteString("\", nonce=\"")
	b.WriteString(wa.Nonce)
	b.WriteString("\"\r\n")
	return b.Bytes()
}

type StateLine struct {
	SipVersion string
	State      *State
}

func (stateline *StateLine) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString(stateline.SipVersion)
	b.WriteByte(' ')
	b.WriteString(stateline.State.Code)
	b.WriteByte(' ')
	b.WriteString(stateline.State.Message)
	b.WriteString("\r\n")
	return b.Bytes()
}

type State struct {
	Code    string
	Message string
}
