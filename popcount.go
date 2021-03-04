// Package popcount defines a function for calculating
// the population count of a uint64 value
package popcount

// Calculate returnes the population count for
// a uint64 value
func Calculate(x uint64) int {
	bits := 0

	for i := 0; i < 64; i++ {
		if x&(1<<i) != 0 {
			bits++
		}
	}

	return bits
}
