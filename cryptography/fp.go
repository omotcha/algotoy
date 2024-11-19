package cryptography

func fpInverse(x, p int) int {
	// recursively calculate inverse of x under prime modulo p
	if x == 0 || x >= p {
		panic("x should be between 1 and p-1")
		return 0
	}
	if x == 1 {
		return 1
	}
	return (p - p/x) * fpInverse(p%x, p) % p
}
