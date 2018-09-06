package gosip

import (
	log "github.com/gogap/logrus"
	"github.com/lastsweetop/gosip/sipHeaderName"
	"strconv"
	"strings"
)

func UnmarshalSipResponse(data []byte, sipresp *SipResponse) {
	content := string(data)
	lines := strings.Split(content, "\r\n\r\n")
	headers := strings.Split(strings.TrimSpace(lines[0]), "\r\n")
	sipresp.StateLine = UnmarshalStateLine(headers[0])
	for _, v := range headers[1:] {
		kv := strings.SplitN(v, ":", 2)
		switch kv[0] {
		case sipHeaderName.Via:
			sipresp.Via = UnmarshalVia(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.From:
			sipresp.From = UnmarshalFrom(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.To:
			sipresp.To = UnmarshalTo(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.CallId:
			sipresp.CallID = UnmarshalCallID(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.CSeq:
			sipresp.Cseq = UnmarshalCSeq(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.ContentLength:
			value, _ := strconv.ParseInt(strings.TrimSpace(kv[1]), 10, 32)
			sipresp.ContentLength = int(value)
		case sipHeaderName.WWWAuthorization:
			sipresp.WWWAuthenticate = UnmarshalWWWAuthorization(strings.TrimSpace(kv[1]))
			break
		}
	}
}

func UnmarshalSipRequest(data []byte, sipreq *SipRequest) {
	content := string(data)
	lines := strings.Split(content, "\r\n\r\n")
	headers := strings.Split(strings.TrimSpace(lines[0]), "\r\n")
	sipreq.StartLine = UnmarshalStartLine(headers[0])

	for _, v := range headers[1:] {
		//log.Println(">>>>>>", v)
		kv := strings.SplitN(v, ":", 2)
		switch kv[0] {
		case sipHeaderName.Via:
			sipreq.Via = UnmarshalVia(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.From:
			sipreq.From = UnmarshalFrom(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.To:
			sipreq.To = UnmarshalTo(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.CallId:
			sipreq.CallID = UnmarshalCallID(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.CSeq:
			sipreq.Cseq = UnmarshalCSeq(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.Contact:
			sipreq.Contact = UnmarshalContact(strings.TrimSpace(kv[1]))
			break
		case sipHeaderName.MaxForwards:
			value, _ := strconv.ParseInt(strings.TrimSpace(kv[1]), 10, 32)
			sipreq.MaxForwards = int(value)
			break
		case sipHeaderName.UserAgent:
			sipreq.UserAgent = strings.TrimSpace(kv[1])
			break
		case sipHeaderName.Expires:
			value, _ := strconv.ParseInt(strings.TrimSpace(kv[1]), 10, 32)
			sipreq.Expires = int(value)
			break
		case sipHeaderName.ContentLength:
			value, _ := strconv.ParseInt(strings.TrimSpace(kv[1]), 10, 32)
			sipreq.ContentLength = int(value)

		case sipHeaderName.WWWAuthorization:
			sipreq.Authorization = UnmarshalAuthorization(strings.TrimSpace(kv[1]))
			break
		}
	}
	if sipreq.ContentLength > 0 {
		sipreq.Body = strings.SplitN(lines[1], "\n", 2)[1]
	}
}

func UnmarshalWWWAuthorization(wwwastr string) *WWWAuthenticate {
	wwwa := &WWWAuthenticate{}
	strs := strings.SplitN(wwwastr, " ", 2)
	log.Println(strs[1])
	params := strings.Split(strs[1], ",")
	for _, v := range params {
		kv := strings.Split(v, "=")
		k := strings.TrimSpace(kv[0])
		v1 := strings.Trim(kv[1], "\"")
		switch k {
		case "realm":
			wwwa.Realm = v1
			break
		case "nonce":
			wwwa.Nonce = v1
			break
		}
	}
	return wwwa
}

func UnmarshalAuthorization(contactstr string) *Authorization {
	auth := &Authorization{}
	strs := strings.SplitN(contactstr, " ", 2)
	log.Println(strs[1])
	params := strings.Split(strs[1], ",")
	for _, v := range params {
		kv := strings.Split(v, "=")
		k := strings.TrimSpace(kv[0])
		v1 := strings.Trim(kv[1], "\"")
		switch k {
		case "username":
			auth.Username = v1
			break
		case "realm":
			auth.Realm = v1
			break
		case "nonce":
			auth.Nonce = v1
			break
		case "uri":
			auth.Uri = v1
			break
		case "response":
			auth.Response = v1
			break
		case "algorithm":
			auth.Algorithm = v1
			break
		}
	}
	log.Println("UnmarshalAuthorization", auth)
	return auth
}

func UnmarshalContact(contactstr string) *Contact {
	contact := &Contact{}
	strs := strings.Split(contactstr, ";")
	addrs := strings.Split(strs[0], "<")
	contact.DisplayName = addrs[0]
	contact.Url = UnmarshalSipUrl(strings.TrimRight(addrs[1], ">"))
	return contact
}

func UnmarshalCSeq(cseqstr string) *Cseq {
	cseq := &Cseq{}
	strs := strings.Split(cseqstr, " ")
	value, _ := strconv.ParseInt(strs[0], 10, 32)
	cseq.Seq = int(value)
	cseq.Method = strs[1]
	return cseq
}

func UnmarshalCallID(callidstr string) *CallID {
	callID := &CallID{}
	strs := strings.Split(callidstr, "@")
	callID.ID = strs[0]
	if len(strs) == 2 {
		callID.Host = strs[1]
	}
	return callID
}

func UnmarshalTo(tostr string) *To {
	to := &To{}
	strs := strings.Split(tostr, ">;")
	addrs := strings.Split(strs[0], "<")
	to.DisplayName = addrs[0]
	to.Url = UnmarshalSipUrl(strings.TrimRight(addrs[1], ">"))
	return to
}

func UnmarshalFrom(fromstr string) *From {
	from := &From{}
	strs := strings.Split(fromstr, ">;")
	addrs := strings.Split(strs[0], "<")
	from.DisplayName = addrs[0]
	from.Url = UnmarshalSipUrl(addrs[1])
	from.Tag = strings.Split(strs[1], "=")[1]

	return from
}

func UnmarshalVia(viastr string) *Via {
	via := &Via{}
	strs := strings.Split(viastr, " ")
	log.Println(strs)
	protocols := strings.Split(strs[0], "/")
	via.SipVersion = protocols[0] + "/" + protocols[1]
	via.ProtocolType = protocols[2]

	viaurls := strings.Split(strs[1], ";")

	addr := strings.Split(viaurls[0], ":")
	via.Address = &Address{addr[0], addr[1]}
	for _, v := range viaurls[1:] {
		kv := strings.Split(v, "=")
		switch kv[0] {
		case "rport":
			if len(kv) == 1 {
				via.Rport = 0
			} else {
				value, _ := strconv.ParseInt(kv[1], 10, 32)
				via.Rport = int(value)
			}
		case "branch":
			via.Branch = kv[1]
			break
		}
	}
	log.Println(via)
	return via
}

func UnmarshalStartLine(linestr string) *StartLine {
	startLine := &StartLine{}
	startLineStrs := strings.Split(linestr, " ")
	startLine.SipMethod = startLineStrs[0]
	startLine.SipUrl = UnmarshalSipUrl(startLineStrs[1])
	startLine.SipVersion = startLineStrs[2]
	return startLine
}

func UnmarshalStateLine(linestr string) *StateLine {
	stateLine := &StateLine{}
	stateLineStrs := strings.Split(linestr, " ")
	stateLine.SipVersion = stateLineStrs[0]
	stateLine.State = &State{stateLineStrs[1], stateLineStrs[2]}
	return stateLine
}

func UnmarshalSipUrl(urlstr string) (sipUrl *SipUrl) {
	sipUrl = &SipUrl{}
	data := strings.TrimLeft(urlstr, "sip:")
	urlstrs := strings.Split(data, ";")
	addresses := strings.Split(urlstrs[0], "@")
	auth := strings.Split(addresses[0], ":")
	url := strings.Split(addresses[1], ":")
	sipUrl.UserName = auth[0]
	sipUrl.Host = url[0]
	if len(auth) > 1 {
		sipUrl.Password = auth[1]
	}
	if len(url) > 1 {
		sipUrl.Port = url[1]
	}
	return
}
