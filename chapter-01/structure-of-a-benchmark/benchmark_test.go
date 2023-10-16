package benchmark

import "testing"

// Simple sequential benchmark.

func BenchmarkFibonacci(b *testing.B) {
	// Setup/iniitlaizatoin.

	// Register cleanup functions.
	b.Cleanup(func() {
		// Code to execute once benchmarking is done.
		//
		// We use .Cleanup() instead of defer because deferred functions
		// aren't guaranteed to execute if the function panics.
		//
		// Cleanup functions are pushed onto a stack and executed in lifo order.
		// similar to deferred funcitons.
	})

	b.ResetTimer() // Rest benchmark timer to ignore execution time prior to inner loop.

	z := 1
	for i := 0; i < b.N; i++ {

		// CODE TO BENCHMARK HERE

		b.StopTimer()
		// Code that should be included in the execution time of the benchmark.
		z = z * (z + 1)
		b.StartTimer()
	}
}
