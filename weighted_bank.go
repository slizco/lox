package lox

// WeightedBank is for storing locks of specified weight
type weightedBank struct {
	m map[string](chan struct{})
	w int
}

// NewWeightedBank returns an empty, initialized bank where each
// resource's lock will have a weight equal to the passed weight.
func NewWeightedBank(w int) weightedBank {
	return weightedBank{
		m: make(map[string]chan struct{}, w),
		w: w,
	}
}

// Lock prevents goroutines from accessing the following code
// when "w" (where "w" equals the resource's lock weight) are
// already in use.
func (b weightedBank) Lock(resource string) {
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
func (b weightedBank) Unlock(resource string) {
	if _, ok := b.m[resource]; ok {
		b.m[resource] <- struct{}{}
	}
}
