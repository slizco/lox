package lox

// Bank is for storing locks of weight 1
type Bank struct {
	m map[string](chan struct{})
}

// NewBank returns an empty, initialized bank where each
// resource's lock will have a weight of 1.
func NewBank() Bank {
	return Bank{make(map[string]chan struct{}, 1)}
}

// Lock prevents any following code from being accessed by
// another goroutine until the resource has been unlocked.
func (b Bank) Lock(resource string) {
	if _, ok := b.m[resource]; ok {
		<-b.m[resource]
	} else {
		b.m[resource] = make(chan struct{}, 1)
	}
}

// Unlock opens up a resource so any code wrapped between
// Lock and Unlock may now be accessed by another goroutine.
func (b Bank) Unlock(resource string) {
	if _, ok := b.m[resource]; ok {
		b.m[resource] <- struct{}{}
	}
}

// WeightedBank is for storing locks of specified weight
type WeightedBank struct {
	m map[string](chan struct{})
	w int
}

// NewWeightedBank returns an empty, initialized bank where each
// resource's lock will have a weight equal to the passed weight.
func NewWeightedBank(w int) WeightedBank {
	return WeightedBank{
		m: make(map[string]chan struct{}, w),
		w: w,
	}
}

// Lock prevents goroutines from accessing the following code
// when "w" (where "w" equals the resource's lock weight) are
// already in use.
func (b WeightedBank) Lock(resource string) {
	if _, ok := b.m[resource]; ok {
		<-b.m[resource]
	} else {
		b.m[resource] = make(chan struct{}, b.w)
		for i := 0; i < b.w-1; i++ {
			b.m[resource] <- struct{}{}
		}
	}
}

// Unlock decrements a resource's currently locked value so any code wrapped
// between Lock and Unlock may now be accessed by another goroutine.
func (b WeightedBank) Unlock(resource string) {
	if _, ok := b.m[resource]; ok {
		b.m[resource] <- struct{}{}
	}
}
