package sip

import (
	"bytes"
	"net"
	"strconv"
)

type SipRequest struct {
	StartLine *StartLine
	SipHeader
	CommonHeader
	Authorization *Authorization
	Body          string
}

type Authorization struct {
	Username  string
	Realm     string
	Nonce     string
	Uri       string
	Response  string
	Algorithm string
}

func (auth Authorization) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("Authorization: Digest ")
	b.WriteString("username=\"")
	b.WriteString(auth.Username)
	b.WriteString("\", ")
	b.WriteString("realm=\"")
	b.WriteString(auth.Realm)
	b.WriteString("\", ")
	b.WriteString("nonce=\"")
	b.WriteString(auth.Nonce)
	b.WriteString("\", ")
	b.WriteString("uri=\"")
	b.WriteString(auth.Uri)
	b.WriteString("\", ")
	b.WriteString("response=\"")
	b.WriteString(auth.Response)
	b.WriteString("\", ")
	b.WriteString("algorithm=")
	b.WriteString(auth.Algorithm)
	b.WriteByte('\n')
	return b.Bytes()
}

type SipHeader struct {
	Via         *Via
	From        *From
	To          *To
	CallID      *CallID
	Cseq        *Cseq
	Contact     *Contact
	MaxForwards int
}

type CommonHeader struct {
	UserAgent     string
	Expires       int
	ContentLength int
}

type Via struct {
	SipVersion   string
	ProtocolType string
	Address      *Address
	Rport        int
	Received     net.IP
	Branch       string
}

func (via Via) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("Via: ")
	b.WriteString(via.SipVersion)
	b.WriteString("/")
	b.WriteString(via.ProtocolType)
	b.WriteString(" ")
	b.WriteString(via.Address.Host)
	b.WriteString(":")
	b.WriteString(via.Address.Port)
	b.WriteString(";rport")
	if via.Rport != 0 {
		b.WriteString("=")
		b.WriteString(strconv.Itoa(via.Rport))
	}
	b.WriteString(";branch=")
	b.WriteString(via.Branch)
	if via.Received != nil {
		b.WriteString(";")
		b.WriteString("received=")
		b.WriteByte(via.Received[0])
		b.WriteByte('.')
		b.WriteByte(via.Received[1])
		b.WriteByte('.')
		b.WriteByte(via.Received[2])
		b.WriteByte('.')
		b.WriteByte(via.Received[3])
	}
	b.WriteString("\r\n")
	return b.Bytes()
}

type Address struct {
	Host string
	Port string
}

type From struct {
	DisplayName string
	Url         *SipUrl
	Tag         string
}

func (from From) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("From: <sip:")
	b.WriteString(from.Url.UserName)
	b.WriteByte('@')
	b.WriteString(from.Url.Host)
	if from.Url.Port != "" {
		b.WriteByte(':')
		b.WriteString(from.Url.Port)
	}
	b.WriteString(">;")
	if from.Tag != "" {
		b.WriteString("tag=")
		b.WriteString(from.Tag)
	}
	b.WriteString("\r\n")
	return b.Bytes()
}

type To struct {
	DisplayName string
	Url         *SipUrl
	Tag         string
}

func (to To) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("To: <sip:")
	b.WriteString(to.Url.UserName)
	b.WriteByte('@')
	b.WriteString(to.Url.Host)
	if to.Url.Port != "" {
		b.WriteByte(':')
		b.WriteString(to.Url.Port)
	}

	b.WriteString(">")
	if to.Tag != "" {
		b.WriteByte(';')
		b.WriteString("tag=")
		b.WriteString(to.Tag)
	}
	b.WriteString("\r\n")
	return b.Bytes()
}

type CallID struct {
	ID   string
	Host string
}

func (callID CallID) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("Call-ID: ")
	b.WriteString(callID.ID)
	b.WriteString("\r\n")
	return b.Bytes()
}

type Cseq struct {
	Seq    int
	Method string
}

func (cseq Cseq) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("CSeq: ")
	b.WriteString(strconv.Itoa(cseq.Seq))
	b.WriteByte(' ')
	b.WriteString(cseq.Method)
	b.WriteString("\r\n")
	return b.Bytes()
}

type Contact struct {
	DisplayName string
	Url         *SipUrl
	Q           float32
	Expires     int
}

func (contact Contact) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("Contact: <sip:")
	b.WriteString(contact.Url.UserName)
	b.WriteByte('@')
	b.WriteString(contact.Url.Host)
	b.WriteByte(':')
	b.WriteString(contact.Url.Port)
	b.WriteString(">\r\n")
	return b.Bytes()
}

type StartLine struct {
	SipMethod  string
	SipUrl     *SipUrl
	SipVersion string
}

func (startLine StartLine) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString(startLine.SipMethod)
	b.WriteString(" sip:")
	b.WriteString(startLine.SipUrl.UserName)
	b.WriteByte('@')
	b.WriteString(startLine.SipUrl.Host)
	b.WriteByte(':')
	b.WriteString(startLine.SipUrl.Port)
	b.WriteByte(' ')
	b.WriteString(startLine.SipVersion)
	b.WriteString("\r\n")
	return b.Bytes()
}

type SipUrl struct {
	UserName  string
	Password  string
	Host      string
	Port      string
	Transport string
	User      string
	Method    string
	HeadName  string
	HeadValue string
}

func (sipUrl SipUrl) BytesBuffer() []byte {
	b := bytes.Buffer{}
	b.WriteString("sip:")
	b.WriteString(sipUrl.UserName)
	b.WriteByte('@')
	b.WriteString(sipUrl.Host)
	b.WriteByte(':')
	b.WriteString(sipUrl.Port)
	return b.Bytes()
}
