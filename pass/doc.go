// Package pass encapsulates breadth-wide search (or, if more precisely, pass)
// through radix tree. Its main purpose is to properly reindex nodes of the
// provided tree to allow enumerate children indices by couples (first, amount)
// or (low, high).
package pass
