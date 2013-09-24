package gangliamr

func nonEmpty(s ...string) string {
	for _, e := range s {
		if e != "" {
			return e
		}
	}
	return ""
}

func makeOptional(base, extra string) string {
	if base == "" {
		return ""
	}
	return base + " " + extra
}
