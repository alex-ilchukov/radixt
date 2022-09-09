package radixt

// Hint on interpretation of amount of bytes, returned by [Hoarder.Hoard]
// method.
const (
	HoardExactly = iota
	HoardAtLeast
)

// Hoader is ancillary interface for radix tree implementations. On realizing
// the interface, the implementations should choose consistent mode accordingly
// to [HoardExactly] and [HoardAtLeast].
type Hoarder interface {
	// Hoard should return amount of bytes, taken by the implementation,
	// with hint how to interpret the amount (see [HoardExactly] and
	// [HoardAtLeast]).
	Hoard() (amount, hint uint)
}
