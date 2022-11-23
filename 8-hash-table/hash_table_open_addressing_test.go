package hashtable

import (
	"testing"
)

func TestOpenAddressing_Set(t *testing.T) {
	ages := NewOpenAddressing(4, NewLinearProber())

	ages.Set("Mohammad", 1)
	ages.Set("Alice", 22)
	ages.Set("Kaz", 32)

	if ages.Size() != 3 {
		t.Fatalf("hash table size expected to be 4, got %d", ages.Size())
	}

	if age := ages.Get("Mohammad"); age != 1 {
		t.Fatalf("entity value expected to be 1, got %d", age)
	}
}

func TestOpenAddressing_Collision(t *testing.T) {
	ages := NewOpenAddressing(4, NewLinearProber())

	ages.Set("Mohammad", 1)
	ages.Set("Alice", 22)
	ages.Set("Kaz", 32)
	ages.Set("Mohammad", 2)
	ages.Set("Kaz", 34)

	if age := ages.Get("Mohammad"); age != 2 {
		t.Fatalf("entity value expected to be 2, got %d", age)
	}

	if age := ages.Get("Kaz"); age != 34 {
		t.Fatalf("entity value expected to be 34, got %d", age)
	}
}

func TestOpenAddressing_Get(t *testing.T) {
	speeds := NewOpenAddressing(10, NewLinearProber())

	speeds.Set("Nissan", 200)
	speeds.Set("Maserati", 280)
	speeds.Set("Ferrari", 320)

	if speed := speeds.Get("Nissan"); speed != 200 {
		t.Fatalf("entity value expected to be 200, got %v", speed)
	}

	if speed := speeds.Get("Maserati"); speed != 280 {
		t.Fatalf("entity value expected to be 280, got %v", speed)
	}

	if speed := speeds.Get("Ferrari"); speed != 320 {
		t.Fatalf("entity value expected to be 320, got %v", speed)
	}

	if speed := speeds.Get("Porsche"); speed != nil {
		t.Fatalf("entity value expected to be nil, got %v", speed)
	}
}

func TestOpenAddressing_Unset(t *testing.T) {
	ages := NewOpenAddressing(2, NewLinearProber())

	ages.Set("Mohammad", 1)
	ages.Set("Alice", 22)
	ages.Set("Kaz", 32)

	ages.Unset("Alice")

	if age := ages.Get("Alice"); age != nil {
		t.Fatalf("entity value expected to be nil, got %v", age)
	}

	if ages.Size() != 2 {
		t.Fatalf("hash table size expected to be 2, got %d", ages.Size())
	}
}
