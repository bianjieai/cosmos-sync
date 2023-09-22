package types

// DenomTrace contains the base denomination for ICS20 fungible tokens and the
// source tracing information path.
type DenomTrace struct {
	// path defines the chain of port/channel identifiers used for tracing the
	// source of the fungible token.
	Path string
	// base denomination of the relayed fungible token.
	BaseDenom string
}
