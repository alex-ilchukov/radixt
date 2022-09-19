package lookup

// Default is used in lookup implementation as a parameter type to set up
// default, non-assisted way of lookup in radix trees.
type Default [0]Switcher

// Switch is used in lookup implementation as a parameter type to set up a way
// of lookup in radix trees, which is assisted (and possible sped up) by
// [Switcher] implementation.
type Switch [1]Switcher
