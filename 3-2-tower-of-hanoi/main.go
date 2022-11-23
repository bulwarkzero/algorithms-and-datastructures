package main

import "fmt"

/*

    _|_ 1            |            |
   __|__ ...         |            |
  ___|___ n-1        |            |
 ____|____ n     ____|____    ____|____
###########   ###########  ###########
 Stack1         Stack2       Stack3

 minimum moves: (2^n) - 1
*/

func towerOfHanoi(n, source, auxiliary, destination int) {
	if n == 1 {
		fmt.Printf("Move disk [%d] from tower [%d] to [%d]\n", n, source, destination)

		return
	}

	// move n-1 disks from source to auxiliary using destination
	towerOfHanoi(n-1, source, destination, auxiliary)

	// move 1 disk (nth item) from source to destination
	towerOfHanoi(1, source, auxiliary, destination)

	// move n-1 disks from auxiliary to destination using source
	towerOfHanoi(n-1, auxiliary, source, destination)
}

func main() {
	// move 3 disks from tower 1 to tower 3 using tower 2
	towerOfHanoi(3, 1, 2, 3)
}
