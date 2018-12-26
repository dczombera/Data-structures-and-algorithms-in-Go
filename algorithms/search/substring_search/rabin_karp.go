package substring_search

import "crypto/rand"

type RabinKarp struct {
	patHash uint64
	Q       uint64
	R       uint64
	M       int
	RM      uint64
}

func RabinKarpConstructor(pattern string, r uint64) RabinKarp {
	p, err := rand.Prime(rand.Reader, 512)
	if err != nil {
		panic(err)
	}
	q := p.Uint64()
	var rm uint64 = 1
	for i := 1; i < len(pattern); i++ {
		rm = (rm * r) % q
	}
	rk := RabinKarp{0, q, r, len(pattern), rm}
	rk.patHash = rk.hash(pattern)
	return rk
}

func (rk RabinKarp) hash(key string) uint64 {
	var h uint64 = 0
	for i := 0; i < rk.M; i++ {
		h = (rk.R*h + uint64(key[i])) % rk.Q
	}
	return h
}

func (rk RabinKarp) Search(text string) int {
	if len(text) < rk.M {
		return len(text)
	}

	txtHash := rk.hash(text)
	if txtHash == rk.patHash {
		return 0
	}
	for i := rk.M; i < len(text); i++ {
		txtHash = (txtHash + rk.Q - uint64(text[i-rk.M])*rk.RM%rk.Q) % rk.Q
		txtHash = (txtHash*rk.R + uint64(text[i])) % rk.Q
		// Use Monte Carlo correctness with a probability of failure of ~ 10^-20 (uint64)
		if txtHash == rk.patHash {
			return i - rk.M + 1
		}
	}
	return len(text)
}
