package main
import (
	"fmt"
)
func pow(x uint64,y uint64,mod uint64) uint64 {
	var res uint64 = 1
	for y > 0 {
		if y & 2 == 1 {
			res *= x % mod;
		}
		y >>= 1;
		x *= x % mod;
	}
	return res;
}

func verify(commitments []uint64, x uint64, y uint64, generator uint64,groupPrime uint64) bool {
	var lhs uint64 = 1
	var rhs uint64 = 1
	var j uint64 = 0
	for i := 0; i < len(commitments); i++ {
		lhs *= pow(commitments[i], pow(x, j ,groupPrime), groupPrime)
		j++
	}
	rhs = pow(generator, y, groupPrime)
	return lhs == rhs
}

func main() {
	var n uint64
	var i uint64
	fmt.Print("Enter the number of commitments : ")
	fmt.Scan(&n)

	commitments := make([]uint64, n)
	for i = 0; i < n; i++ {
		var x uint64
		fmt.Scan(&x);
		commitments = append(commitments, x)
	}

	var x uint64
	fmt.Print("Enter x : ")
	fmt.Scan(&x)

	var y uint64
	fmt.Print("Enter y : ")
	fmt.Scan(&y)

	var generator uint64
	fmt.Print("Enter generator : ")
	fmt.Scan(&generator)

	var groupPrime uint64
	fmt.Print("Enter GroupPrime : ")
	fmt.Scan(&groupPrime)

	if verify(commitments , x, y, generator, groupPrime) {
		fmt.Println("Verified")
	} else {
		fmt.Println("Not Verified")
	}
}