package map_ext

import (
	"testing"
)

type discountStrategy interface {
	Discount(p float32) float32
}

type normalDiscount struct {
}

func (d *normalDiscount) Discount(p float32) float32 {
	return p * 0.9
}

type double11Discount struct {
}

func (d *double11Discount) Discount(p float32) float32 {
	return p * 0.5
}

func TestMapWithStrategyPattern(t *testing.T) {
	m := map[string]discountStrategy{
		"normal":   &normalDiscount{},
		"double11": &double11Discount{},
	}
	t.Log(m["normal"].Discount(100), m["double11"].Discount(100)) // 90 50
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
