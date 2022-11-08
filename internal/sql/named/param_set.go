package named

// ParamSet represents a set of parameters for a single query
type ParamSet struct {
	// does this engine support named parameters?
	hasNamedSupport bool
	// the set of 