package go_evm384

import (
	"math/bits"
)

const NUM_LIMBS = 6
type Element [NUM_LIMBS]uint64

// Cmp compares u and v and returns:
//
//   -1 if u <  v
//    0 if u == v
//   +1 if u >  v
//
func cmp128(u_hi, u_lo, v_hi, v_lo uint64) int {
	if u_hi == v_hi && u_lo == v_lo {
		return 0
	} else if u_hi < v_hi || (u_hi == v_hi && u_lo < v_lo) {
		return -1
	} else {
		return 1
	}
}

// Add64 returns u+v.
func Add64(u_hi, u_lo, v_hi, v_lo uint64) (uint64, uint64) {
	lo, carry := bits.Add64(u_lo, v_lo, 0)
	hi := u_hi + carry
	return hi, lo
}


// out <- x + y
func Add(out *Element, x *Element, y *Element) {
	var c uint64
	c = 0

	// TODO: manually unroll this?
	for i := 0; i < NUM_LIMBS; i++ {
		// out[i] = x[i] + y[i] + c
		out[i], c = bits.Add64(x[i], y[i], c)
	}
}

// out <- x - y
func Sub(out *Element, x *Element, y *Element) (uint64){
	var c uint64
	c = 0

	for i := 0; i < NUM_LIMBS; i++ {
		out[i], c = bits.Sub64(x[i], y[i], c)
	}

	return c
}

// return x <= y
func lte(x *Element, y *Element) bool {
	for i := 0; i < NUM_LIMBS; i++ {
		if x[i] > y[i] {
			return false
		}
	}

	return true
}

/*
	Modular Addition
*/
func AddMod(out *Element, x *Element, y *Element, mod *Element) {
	Add(out, x, y)

	if lte(mod, out) {
		Sub(out, out, mod)
	}
}

/*
	Modular Subtraction
*/
func SubMod(out *Element, x *Element, y *Element, mod *Element) {
	var c uint64
	c = Sub(out, x, y)

	// if result < 0 -> result += mod
	if c != 0 {
		Add(out, out, mod)
	}
}

/*
	Montgomery Modular Multiplication: algorithm 14.36, Handbook of Applied Cryptography, http://cacr.uwaterloo.ca/hac/about/chap14.pdf
*/
func MulMod(out *Element, x *Element, y *Element, mod *Element, inv uint64) {
	var A [NUM_LIMBS * 2 + 1]uint64
	var xiyj_hi, xiyj_lo, uimj_hi, uimj_lo, partial_sum_hi, partial_sum_lo, sum_hi, sum_lo uint64 = 0, 0, 0, 0, 0, 0, 0, 0
        var ui, carry uint64 = 0, 0

	for i := 0; i < NUM_LIMBS; i++ {
		ui = (A[i] + x[i] * y[0]) * inv

		carry = 0
		for j := 0; j < NUM_LIMBS; j++ {
			xiyj_lo = x[i]
			xiyj_hi, xiyj_lo = bits.Mul64(xiyj_lo, y[j])

			uimj_lo = ui
			uimj_hi = 0

			uimj_hi, uimj_lo = bits.Mul64(uimj_lo, mod[j])
			partial_sum_hi, partial_sum_lo = Add64(xiyj_hi, xiyj_lo, 0, carry)

			/*
			sum = uimj.Add64(A[i + j])
			sum = sum.Add(partial_sum)
			*/

			sum_hi, sum_lo = Add64(uimj_hi, uimj_lo, 0, A[i + j])
			sum_hi, sum_lo = Add64(sum_hi, sum_lo, partial_sum_hi, partial_sum_lo)

			A[i + j] = sum_lo
			carry = sum_hi

			if cmp128(sum_hi, sum_lo, partial_sum_hi, partial_sum_lo) == -1 {
				var k int
				k = 2
				for ; i + j + k < NUM_LIMBS * 2 && A[i + j + k] == ^uint64(0); {
					A[i + j + k] = 0
					k++
				}

				if (i + j + k < NUM_LIMBS * 2 + 1) {
					A[i + j + k] += 1
				}
			}
		}

		A[i + NUM_LIMBS] += carry
	}

	for i := 0; i < NUM_LIMBS; i++ {
		out[i] = A[i + NUM_LIMBS]
	}

	if lte(mod, out) {
		Sub(out, out, mod)
	}
}
