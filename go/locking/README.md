## TO RUN BENCHMARK 


### BENCHMARK COMPARE
<code>
go test -bench=BenchmarkCounterCompare -benchmem -count=10 locking/mutex.go locking/atomic.go locking/mutex_test.go > locking/compare.txt
</code>


### BENCHMARK ATOMIC, MUTEX
<code>
go test -bench=. -count=10 -benchmem locking/atomic.go locking/atomic_test.go > locking/atomic.txt
go test -bench=. -count=10 -benchmem locking/mutex.go locking/mutex_test.go > locking/mutex.txt
</code>


### BENCHSTAT for statistic
<code>
benchstat -row /Size -col /Method locking/compare.txt
benchstat -col /Method locking/compare.txt
</code>