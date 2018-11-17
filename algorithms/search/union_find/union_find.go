package union_find

// UF data type models connectivity among a set of n sites.
// This implementation uses weighted quick union by rank with path compression by halving during find operation.
type UF struct {
	parent []int
	rank   []byte
	count  int
}

func NewUnionFind(n int) *UF {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UF{parent, make([]byte, n), n}
}

func (uf *UF) Find(p int) int {
	uf.validateSite(p)
	for p != uf.parent[p] {
		// Path compression by halving
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}

func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP != rootQ {
		// Always connect smaller tree to larger one
		if uf.rank[rootP] < uf.rank[rootQ] {
			uf.parent[rootP] = rootQ
		} else if uf.rank[rootQ] < uf.rank[rootP] {
			uf.parent[rootQ] = rootP
		} else {
			uf.parent[rootP] = rootQ
			uf.rank[rootQ]++
		}
		uf.count--
	}
}

func (uf *UF) Connected(p, q int) bool {
	uf.validateSites(p, q)
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) validateSites(pp ...int) {
	for _, p := range pp {
		uf.validateSite(p)
	}
}

func (uf *UF) validateSite(p int) {
	if p < 0 || p >= len(uf.parent) {
		panic("Site out of bounds")
	}
}
