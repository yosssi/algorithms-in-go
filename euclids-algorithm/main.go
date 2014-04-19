// Euclid's algorithm
//
// Compute the greatest common divisor of
// two nonnegative integers p and q as follows:
// If q is 0, the answer is p. If not, divide p by q
// and take the remainder r. The answer is the
// greatest common divisor of q and r.

package main

import (
	"fmt"
	"math"
)

func main() {
	execGCD(2, 4)
	execGCD(3, 9)
	execGCD(4, 6)
	execGCD(4, 8)
	execGCD(12, 22)
	execGCD(12, 42)
	execGCD(12, 48)
}

func execGCD(p int, q int) {
	fmt.Printf("p: %d, q: %d, gcd: %d\n", p, q, gcd(p, q))
}

func gcd(p int, q int) int {
	if q == 0 {
		return p
	}
	r := int(math.Mod(float64(p), float64(q)))
	return gcd(q, r)
}
