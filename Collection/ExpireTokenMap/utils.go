package expiretokenmap

type Options struct {
	AdaptiveCleanInterval bool // Dynamic adjust clean interval to keep memory usage
	TargetLen             int  // Target length with adaptive cleaning
}
