package bench

func CloneList(in []string) []string {
	// nolint prealloc
	var out []string

	return append(out, in...)
}
