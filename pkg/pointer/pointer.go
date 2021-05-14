package pointer

func CopyValue(p *int) *int {
	var value *int

	*value = *p
	return value
}

func CopyRef(p *int) *int {
	var value *int

	value = p
	return value
}

func Me(p *int) *int {
	return p
}
