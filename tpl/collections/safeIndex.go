package collections

// SafeIndex returns the result of indexing its first argument by the following
// arguments. Thus "index x 1 2 3" is, in Go syntax, x[1][2][3]. Each
// indexed item must be a map, slice, or array. Doesn't return errors, just empty value.
func (ns *Namespace) SafeIndex(item interface{}, indexes ...interface{}) interface{} {
	item, err := ns.Index(item, indexes...)
	if err != nil {
		return nil
	}

	return item
}
