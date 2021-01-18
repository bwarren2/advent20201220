package advent20201220

type Edge string

func (e Edge) AsCanonical() Edge {
	if e > e.Reverse() {
		return e
	}
	return e.Reverse()
}

func (e Edge) Reverse() Edge {
	result := ""
	for _, v := range e {
		result = string(v) + result
	}
	return Edge(result)
}
