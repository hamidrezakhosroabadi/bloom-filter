package bloom

type Bloom[T any] struct {
	Bits          []bool
	HashFunctions []func(T) int
}

func (b *Bloom[T]) Push(item T) {
	for _, hash := range b.HashFunctions {
		b.Bits[hash(item)%len(b.Bits)] = true
	}
}

func (b *Bloom[T]) Contains(item T) bool {
	for _, hash := range b.HashFunctions {
		if !b.Bits[hash(item)%len(b.Bits)] {
			return false
		}
	}
	return true
}

func (b *Bloom[T]) Clear() {
	b.Bits = make([]bool, len(b.Bits))
}

func NewBloom[T any](size int, hashFunctions ...func(T) int) Bloom[T] {
	return Bloom[T]{
		Bits:          make([]bool, size),
		HashFunctions: hashFunctions,
	}
}
