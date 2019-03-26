package main

func main() {
	sum := 1
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }

	// without pre/post statements
	for sum < 1000 {
		sum += sum
	}

	// as a while loop
	for sum < 1000 {
		sum += sum
	}

	// infinite loop
	for {
		// do something in a loop forever
	}

}
