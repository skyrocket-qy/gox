TL;DR: When to use what?
Problem	Use This Profile
App is slow	CPU
High memory usage / leaks	Heap, allocs
High latency on goroutines	Block
Sync locks feel slow	Mutex
Lots of goroutines?	Goroutine
Too many threads?	Threadcreate
Full performance visibility	Trace