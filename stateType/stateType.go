package stateType

import "github.com/lastsweetop/gosip"

var (
	Unauthorized gosip.State = gosip.State{"401", "Unauthorized"}
	OK           gosip.State = gosip.State{"200", "OK"}
	Forbidden    gosip.State = gosip.State{"403", "Forbidden"}
)
