package main

import (
	"fmt"
)
//todo : res should be replaced with different data type
func pow(x uint64, y uint64, mod uint64) uint64 {
	var res uint64 = 1
	for y > 0 {
		if y&1 == 1 {
			res = ((res % mod) * (x % mod)) % mod
		}
		y >>= 1
		x = ((x % mod) * (x % mod)) % mod
	}
	return res
}

func verify(commitments []uint64, x uint64, y uint64, generator uint64, fieldPrime uint64, groupPrime uint64) bool {
	var lhs uint64 = 1
	var rhs uint64
	var j uint64 = 0
	for i := 0; i < len(commitments); i++ {
		lhs = ((lhs % groupPrime) * (pow(commitments[i], pow(x, j, fieldPrime), groupPrime) % groupPrime)) % groupPrime
		j++
	}
	rhs = pow(generator, y, groupPrime)
	return lhs == rhs
}

func main() {
	var n int
	fmt.Print("Enter the number of commitments : ")
	fmt.Scan(&n)

	commitments := make([]uint64, n)
	fmt.Printf("Enter %d commitments : ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&commitments[i])
	}

	var m int
	fmt.Print("Enter the number of shares : ")
	fmt.Scan(&m)

	x := make([]uint64, m)
	y := make([]uint64, m)
	for i := 0; i < m; i++ {
		fmt.Printf("Enter x%d : ", i)
		fmt.Scan(&x[i])
		fmt.Printf("Enter y%d : ", i)
		fmt.Scan(&y[i])
	}

	var generator uint64
	fmt.Print("Enter generator : ")
	fmt.Scan(&generator)

	var groupPrime uint64
	fmt.Print("Enter GroupPrime : ")
	fmt.Scan(&groupPrime)

	var fieldPrime uint64
	fmt.Print("Enter FeildPrime : ")
	fmt.Scan(&fieldPrime)

	for i := 0; i < m; i++ {
		if verify(commitments, x[i], y[i], generator, fieldPrime, groupPrime) {
			fmt.Printf("%d Verified", i)
		} else {
			fmt.Printf("%d Not Verified", i)
		}
	}
}
