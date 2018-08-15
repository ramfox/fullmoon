package state

// new type moon

type MoonPhase int

const (
	// different moon phases
	WaningGibbous MoonPhase = iota
	FirstQuarter
	WaningCresent
	NewMoon
	WaxingCresent
	LastQuarter
	WaxingGibbous
	FullMoon
)

// String implements stringer interface for DataFormat
func (m MoonPhase) String() string {
	s, ok := map[MoonPhase]string{
		WaningGibbous: "🌖",
		FirstQuarter:  "🌓",
		WaningCresent: "🌘",
		NewMoon:       "🌑",
		WaxingCresent: "🌒",
		LastQuarter:   "🌗",
		WaxingGibbous: "🌔",
		FullMoon:      "🌕",
	}[m]

	if !ok {
		return ""
	}

	return s
}
