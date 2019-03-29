package lox

// Bank is for storing locks of weight 1
type bank struct {
	m map[string](chan struct{})
}

// NewBank returns an empty, initialized bank where each
// resource's lock will have a weight of 1.
func NewBank() bank {
	return bank{make(map[string]chan struct{}, 1)}
}

// Lock prevents any following code from being accessed by
// another goroutine until the resource has been unlocked.
func (b bank) Lock(resource string) {
	if _, ok := b.m[resource]; ok {
		<-b.m[resource]
	} else {
		b.m[resource] = make(chan struct{}, 1)
	}
}

// Unlock opens up a resource so any code wrapped between
// Lock and Unlock may now be accessed by another goroutine.
func (b bank) Unlock(resource string) {
	if _, ok := b.m[resource]; ok {
		b.m[resource] <- struct{}{}
	}
}
