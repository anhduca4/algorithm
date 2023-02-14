package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func main() {
	b := []int{2, 8, 3, 4, 5, 32}
	for _, c := range b {
		fmt.Println("Check number ", c, isPowerOfTwo(c))
	}
}

// const isPowerOfTwo = (n) => {
//   if (n <= 0) {
//     return false;
//   }
//   return ~(n & n - 1);
// }

// const a = [2, 8, 3, 4, 5, 32];
// for (const i of a) {
//   console.log(i + 'is power of two : ' + isPowerOfTwo(2));
// }
