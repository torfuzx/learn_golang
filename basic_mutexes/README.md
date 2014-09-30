An operation acting on shared memory is atommic if it completes in a single
step relative to other threads. When an atomic store is performed on a shared
memory, no other thread can observe the modification half-complete. When an
atomic load is performed on a shared variable, it reads the entire value as
it appeared at a single moment in time. Non-atomic loads and stores do not make
those guarantees.

Without those guarantees, lock-free programming would be impossible, since you
could never let different threads manipulate a shared variable at the same time.

Any time two threads operate on a shared variable concurrently, and one of
those operations performs a write, both threads must use atomic operations.

A memory operation can be non-atomic because it uses multiple CPU instructions,
non-atomic even when using a single CPU instruction, or non-atomic because you
are writing portable code and you simply can make the assumption.

### References:

- [Atomic vs. Non-Atomic Operations](http://preshing.com/20130618/atomic-vs-non-atomic-operations/)
- [Linearizability](http://en.wikipedia.org/wiki/Linearizability)
