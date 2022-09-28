package analysis

// Default represents default mode of analysis of tree nodes' chunks, when the
// chunks are analyzed and processed as is.
type Default [0]byte

// Firstless represents advanced mode of analysis of tree nodes' chunks, when
// the chunks are analyzed and processed _without their first bytes_. The mode
// concerns the following areas of analysis report:
//
//   - [A.C] field would hold "crammed" chunks without their first bytes;
//   - [N.Chunk] field would hold the node's chunk without its first byte;
//   - [N.ChunkPos] field would hold position of [N.Chunk] within of [A.C].
//
// The following fields would hold the same values as with [Default] mode:
//
//   - [N.ChunkFirst] would hold the first byte of _original_ node's chunk;
//   - [N.ChunkEmpty] would hold the boolean flag on if the node has empty
//     original chunk or not.
type Firstless [1]byte

// Mode is set of types of tree analysis regarding processing of its nodes'
// chunks.
type Mode interface {
	Default | Firstless
}
