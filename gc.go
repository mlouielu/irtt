package irtt

import (
	"fmt"
)

// GCMode selects when the garbage collector is run.
type GCMode int

// StampAt constants.
const (
	GCOn GCMode = iota
	GCOff
	GCIdle
)

var gcms = [...]string{"on", "off", "idle"}

func (gm GCMode) String() string {
	if int(gm) < 0 || int(gm) >= len(gcms) {
		return fmt.Sprintf("GCMode:%d", gm)
	}
	return gcms[gm]
}

// ParseGCMode returns a GCMode value from its string.
func ParseGCMode(s string) (GCMode, error) {
	for i, v := range gcms {
		if v == s {
			return GCMode(i), nil
		}
	}
	return GCOn, Errorf(InvalidGCModeString, "invalid GC mode string: %s", s)
}
