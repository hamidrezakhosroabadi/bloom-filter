package bloom

import "testing"

func hashFunction(input int) int {
	return input
}
func TestPush(t *testing.T) {

	var b Bloom[int] = NewBloom(100, hashFunction)
	b.Push(100)
	b.Push(999)

	if !b.Contains(999) {
		t.Errorf("Expected true, got false")
	}

	if !b.Contains(100) {
		t.Errorf("Expected true, got false")
	}

}

func TestContains(t *testing.T) {

	var b Bloom[int] = NewBloom(100, hashFunction)

	b.Push(100)
	b.Push(999)

	if !b.Contains(999) {
		t.Errorf("Expected true, got false")
	}

	if b.Contains(101) {
		t.Errorf("Expected false, got true")
	}

}
