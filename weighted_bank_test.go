package lox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightedLock(t *testing.T) {
	b := NewWeightedBank(3)
	b.Lock("resource")

	if _, ok := b.m["resource"]; !ok {
		assert.Fail(t, "Failed to lock resource")
	}
	for i := 0; i < 3; i++ {
		select {
		case <-b.m["resource"]:
			if i == 2 {
				assert.FailNow(t, "Locked with too much weight")
			}
			continue
		default:
			if i != 2 {
				assert.FailNow(t, "Locked with not enough weight")
			}
		}
	}
}

func TestWeightedUnLock(t *testing.T) {
	b := NewWeightedBank(3)
	b.m["resource"] = make(chan struct{}, 3)
	b.Unlock("resource")

	select {
	case <-b.m["resource"]:
		return
	default:
		assert.FailNow(t, "Failed to unlock resource")
	}
}
