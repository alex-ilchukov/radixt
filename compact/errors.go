package compact

import "errors"

// ErrorInvalidLenNode is used by common machinery of the compact
// implementations of radix trees to indicate, that there is disrepancy between
// provided bits length and actual bits length of tree node: the former can be
// more than the latter.
var ErrorInvalidLenNode = errors.New(
	"provided bits length of node is more than actual bits length of node",
)

// ErrorOverflow is used by common machinery of the compact implementations of
// radix trees to indicate, that analysis of the provided tree has shown there
// is no possibility to fit node fields into tight bit string.
var ErrorOverflow = errors.New("required fields would not fit into node")
