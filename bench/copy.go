package bench

func CloneList(in []string) []string {
	var out []string

	for _, v := range in {
		out = append(out, v)
	}

	return out
}
