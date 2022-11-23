package hashtable

type linearProber struct {
	linearConst int
}

func (prober *linearProber) Probe(key string, x int) int {
	return x * prober.linearConst
}

func (prober *linearProber) AdjustCapacity(capacity int) int {
	for gcd(prober.linearConst, capacity) != 1 {
		capacity++
	}

	return capacity
}

func NewLinearProber() Prober {
	return &linearProber{17}
}
