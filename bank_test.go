package lox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	b := NewBank()
	b.Lock("resource")

	if _, ok := b.m["resource"]; !ok {
		assert.Fail(t, "Failed to lock resource")
	}

}

func TestUnLock(t *testing.T) {
	b := NewBank()
	b.m["resource"] = make(chan struct{}, 1)
	b.Unlock("resource")

	select {
	case <-b.m["resource"]:
		return
	default:
		assert.Fail(t, "Failed to unlock resource")
	}
}
