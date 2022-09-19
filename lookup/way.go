package lookup

// Default is used in lookup implementation as a parameter type to set up
// default, non-assisted way of lookup in radix trees.
type Default [0]Switcher

// Switch is used in lookup implementation as a parameter type to set up a way
// of lookup in radix trees, which is assisted (and possible sped up) by
// [Switcher] implementation.
type Switch [1]Switcher

// Way is type set of types of lookup ways.
type Way interface {
	Default | Switch
}

// IsSwitch returns, if the provided type parameter from [Way] is [Switch].
func IsSwitch[W Way]() bool {
	var w W
	return len(w) > 0
}
