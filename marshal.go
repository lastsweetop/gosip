package sip

import (
	"bytes"
	"strconv"
)

func MarshalSipResponse(sipresp *SipResponse) (data []byte) {
	b := bytes.Buffer{}

	b.Write(sipresp.StateLine.BytesBuffer())
	b.Write(sipresp.Via.BytesBuffer())
	b.Write(sipresp.From.BytesBuffer())
	b.Write(sipresp.To.BytesBuffer())
	b.Write(sipresp.CallID.BytesBuffer())
	b.Write(sipresp.Contact.BytesBuffer())

	if sipresp.Contact != nil {
		b.Write(sipresp.Contact.BytesBuffer())
	}

	if sipresp.Date != "" {
		b.WriteString("Date: ")
		b.WriteString(sipresp.Date)
		b.WriteString("\r\n")
	}

	if sipresp.WWWAuthenticate != nil {
		b.Write(sipresp.WWWAuthenticate.BytesBuffer())
	}

	b.WriteString("Content-Length: 0")
	b.WriteString("\r\n")
	b.WriteString("\r\n")

	return b.Bytes()
}

func MarshalSipRequest(sipreq *SipRequest) (data []byte) {
	b := bytes.Buffer{}
	b.Write(sipreq.StartLine.BytesBuffer())
	b.Write(sipreq.Via.BytesBuffer())
	b.Write(sipreq.From.BytesBuffer())
	b.Write(sipreq.To.BytesBuffer())
	b.Write(sipreq.CallID.BytesBuffer())
	b.Write(sipreq.Cseq.BytesBuffer())
	b.Write(sipreq.Contact.BytesBuffer())

	if sipreq.Authorization != nil {
		b.Write(sipreq.Authorization.BytesBuffer())
	}

	b.WriteString("Max-Forwards: ")
	b.WriteString(strconv.Itoa(sipreq.MaxForwards))
	b.WriteString("\r\n")
	b.WriteString("User-Agent: ")
	b.WriteString(sipreq.UserAgent)
	b.WriteString("\r\n")
	b.WriteString("Expires: ")
	b.WriteString(strconv.Itoa(sipreq.Expires))
	b.WriteString("\r\n")
	b.WriteString("Content-Length: 0")
	b.WriteString("\r\n")
	b.WriteString("\r\n")

	return b.Bytes()
}
